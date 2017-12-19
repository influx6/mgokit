package trees

import "strings"

// Query defines a package level variable for access the query interface
// which handles running css queries on markup structures.
var Query queryCtrl

type queryCtrl struct{}

// Selector defines a structure which defines the requirements for a given
// matching to be processed.
type Selector struct {
	Tag       string
	ID        string
	Psuedo    string
	AttrOp    string
	AttrName  string
	AttrValue string
	Classes   []string
	Children  []*Selector
	Order     map[string]string
}

// GetSelector returns the selector received for the given selector.
func (s *Selector) GetSelector() string {
	sel := s.Tag

	if s.ID != "" {
		sel += s.GetID()
	}

	if s.Classes != nil {
		sel += s.GetClass()
	}

	if s.AttrName != "" {
		sel += "[" + s.AttrName + s.AttrOp + s.AttrValue + "]"
	}

	if s.Psuedo != "" {
		sel += s.Psuedo
	}

	return sel
}

// GetID returns the id string for this selector.
func (s *Selector) GetID() string {
	if s.ID != "" {
		return "#" + s.ID
	}

	return ""
}

// GetClass returns the class string for this selector.
func (s *Selector) GetClass() string {
	if s.Classes == nil {
		return ""
	}

	var sels []string

	for _, class := range s.Classes {
		sels = append(sels, "."+class)
	}

	return strings.Join(sels, "")
}

// Query returns the first element matching the giving selector.
func (q queryCtrl) Query(root *Markup, sel string) *Markup {
	sels := q.ParseSelector(sel)
	if sels == nil {
		return nil
	}

	return q.QuerySelector(root, sels[0])
}

// QueryAll returns the first element matching the giving selector.
func (q queryCtrl) QueryAll(root *Markup, sel string) []*Markup {
	sels := q.ParseSelector(sel)
	if sels == nil {
		return nil
	}

	return q.QueryAllSelector(root, sels[0])
}

// QuerySelector uses the provided selector and root returning the first
// element that matches the selector's criteria.
func (q queryCtrl) QuerySelector(root *Markup, sel *Selector) *Markup {
	var filtered *Markup

childloop:
	for _, child := range root.children {
		if q.queryOne(child, sel) {
			filtered = child
			break childloop
		}

		for _, kid := range child.children {
			if q.queryOne(kid, sel) {
				filtered = kid
				break childloop
			}

			if item := q.QuerySelector(kid, sel); item != nil {
				filtered = item
				break childloop
			}
		}
	}

	if sel.Children == nil {
		return filtered
	}

	for _, child := range sel.Children {
		filtered = q.QuerySelector(filtered, child)
		if filtered == nil {
			return nil
		}
	}

	return filtered
}

// QueryAllSelector uses the provided selector and root returning all
// elements that matches the selector's criteria.
func (q queryCtrl) QueryAllSelector(root *Markup, sel *Selector) []*Markup {
	var found []*Markup

	for _, child := range root.children {
		if !q.queryOne(child, sel) {

			for _, kid := range child.children {
				if q.queryOne(kid, sel) {
					found = append(found, kid)
				}

				found = append(found, q.QueryAllSelector(kid, sel)...)
			}

			continue
		}

		if sel.Children == nil {
			found = append(found, child)

			for _, kid := range child.children {
				if q.queryOne(kid, sel) {
					found = append(found, kid)
				}

				found = append(found, q.QueryAllSelector(kid, sel)...)
			}

			continue
		}

		kid := child
		for _, kidSel := range sel.Children {
			kid = q.QuerySelector(kid, kidSel)
			if kid == nil {
				continue
			}
		}

		for _, mkid := range child.children {
			if q.queryOne(mkid, sel) {
				found = append(found, mkid)
			}

			found = append(found, q.QueryAllSelector(mkid, sel)...)
		}

		found = append(found, kid)
	}

	return found
}

func (q queryCtrl) queryOne(target *Markup, sel *Selector) bool {
	if sel.Tag != "" && !q.tagFor(target, sel.Tag) {
		return false
	}

	if sel.ID != "" && !q.idFor(target, sel.ID) {
		return false
	}

	if sel.Classes != nil {
		for _, class := range sel.Classes {
			if !q.classFor(target, class) {
				return false
			}
		}

		return true
	}

	if sel.AttrName != "" && !q.attrFor(target, sel.AttrName, sel.AttrValue, sel.AttrOp) {
		return false
	}

	return true
}

var (
	dot         = byte('.')
	coma        = byte(',')
	hash        = byte('#')
	space       = byte(' ')
	bracket     = byte('[')
	endbracket  = byte(']')
	orderOpen   = byte('(')
	orderClosed = byte(')')
)

// ParseSelector returns the giving selector parsed out into its individual
// sections.
func (q queryCtrl) ParseSelector(sel string) []*Selector {
	var sels []*Selector
	items := []byte(sel)
	itemsLen := len(items)

	var index int
	var doChildren bool
	var seenSpace bool
	var seenComa bool

	var child *Selector
	sels = append(sels, &Selector{})

	{

	parseLoop:
		for {
			if index >= itemsLen {
				break
			}

			item := items[index]
			if seenSpace && item == space {
				index++
				continue
			}

			if seenSpace && item != space {
				seenSpace = false
			}

			csel := sels[len(sels)-1]

			switch item {
			case space:
				if seenComa {
					doChildren = false
				}

				if !seenComa {
					doChildren = true
				}

				if !seenComa {
					doChildren = true
				}

				if !seenSpace && !seenComa {
					csel.Children = append(csel.Children, &Selector{})
				}

				index++
				seenSpace = true
				continue parseLoop

			case orderOpen:
				var order []byte
				order = append(order, item)

				{
				nOrderLoop:
					for {
						index++

						if index >= itemsLen {
							ordered := string(order)
							ordered = strings.TrimPrefix(ordered, "(")
							ordered = strings.TrimSuffix(ordered, ")")

							if doChildren {
								if child.Order == nil {
									child.Order = make(map[string]string)
								}
							} else {
								if csel.Order == nil {
									csel.Order = make(map[string]string)
								}
							}

							for _, or := range strings.Split(ordered, ",") {
								ors := strings.Split(or, ":")
								if len(ors) < 2 {
									continue
								}

								if doChildren {
									child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
								} else {
									csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
								}
							}

							break parseLoop
						}

						item = items[index]
						switch item {
						case orderClosed:
							order = append(order, item)

							ordered := string(order)
							ordered = strings.TrimPrefix(ordered, "(")
							ordered = strings.TrimSuffix(ordered, ")")

							if doChildren {
								if child.Order == nil {
									child.Order = make(map[string]string)
								}
							} else {
								if csel.Order == nil {
									csel.Order = make(map[string]string)
								}
							}

							for _, or := range strings.Split(ordered, ",") {
								ors := strings.Split(or, ":")
								if len(ors) < 2 {
									continue
								}

								if doChildren {
									child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
								} else {
									csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
								}
							}

							continue nOrderLoop
						default:
							order = append(order, item)
							continue nOrderLoop
						}
					}
				}

			case coma:
				seenComa = true

				sels = append(sels, &Selector{})

				index++
				continue parseLoop

			case dot:
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
				}

				{
					var blk []byte

				dotLoop:
					for {
						index++

						if index >= itemsLen {
							if len(blk) != 0 {
								if doChildren {
									child.Classes = append(child.Classes, string(blk))
								} else {
									csel.Classes = append(csel.Classes, string(blk))
								}

								blk = nil
							}

							break dotLoop
						}

						item = items[index]

						switch item {
						case dot:
							if doChildren {
								child.Classes = append(child.Classes, string(blk))
							} else {
								csel.Classes = append(csel.Classes, string(blk))
							}

							blk = nil
							continue dotLoop
						case space, coma, hash:
							if doChildren {
								child.Classes = append(child.Classes, string(blk))
							} else {
								csel.Classes = append(csel.Classes, string(blk))
							}

							blk = nil
							continue parseLoop
						case orderOpen:
							var order []byte
							order = append(order, item)

							{
							orderLoop:
								for {
									index++

									item = items[index]
									// fmt.Printf("order: %+q : %+q -> %d : %d\n", order, item, index, itemsLen)

									if index >= itemsLen {
										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										for _, or := range strings.Split(ordered, ",") {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue dotLoop
									}

									switch item {
									case orderClosed:
										order = append(order, item)

										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										for _, or := range strings.Split(ordered, ",") {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue dotLoop
									default:
										order = append(order, item)
										continue orderLoop
									}

								}
							}
						}

						blk = append(blk, item)
					}
				}

				continue parseLoop

			case hash:
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
				}

				{
					var blk []byte

				hashLoop:
					for {
						index++

						if index >= itemsLen {
							if doChildren {
								child.ID = string(blk)
							} else {
								csel.ID = string(blk)
							}

							blk = nil
							break hashLoop
						}

						item = items[index]

						switch item {
						case orderOpen:

							var order []byte
							order = append(order, item)

							{
							hashOrderLoop:
								for {
									index++

									item = items[index]
									if index >= itemsLen {
										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										for _, or := range strings.Split(ordered, ",") {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue hashLoop
									}

									switch item {
									case orderClosed:
										order = append(order, item)

										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										for _, or := range strings.Split(ordered, ",") {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue hashLoop
									default:
										order = append(order, item)
										continue hashOrderLoop
									}

								}
							}

						case dot, space, coma:
							if doChildren {
								child.ID = string(blk)
							} else {
								csel.ID = string(blk)
							}

							blk = nil
							continue parseLoop

						}

						blk = append(blk, item)
					}
				}

			case bracket:
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
				}

				{
					var blk []byte
					blk = append(blk, item)

					for {
						index++

						if index >= itemsLen {
							break
						}

						item = items[index]
						if item == endbracket {
							blk = append(blk, item)

							attr, val, op := q.splitBracketSelector(string(blk))
							val = strings.Replace(val, "'", "", -1)
							val = strings.Replace(val, "\"", "", -1)

							if doChildren {
								child.AttrOp = op
								child.AttrName = attr
								child.AttrValue = val
							} else {
								csel.AttrOp = op
								csel.AttrName = attr
								csel.AttrValue = val
							}

							index++
							continue parseLoop
						}

						blk = append(blk, item)
					}

					continue parseLoop
				}

			default:
				if doChildren {
					child = csel.Children[len(csel.Children)-1]
				}

				{
					var blk []byte
					blk = append(blk, item)

				defaultLoop:
					for {
						index++

						if index >= itemsLen {
							if len(blk) != 0 {
								if doChildren {
									if child.Tag != "" {
										break parseLoop
									}

									child.Tag = string(blk)
									if psud := strings.Index(child.Tag, ":"); psud != -1 {
										psuedo := child.Tag[psud:]
										child.Tag = child.Tag[:psud]
										child.Psuedo = psuedo
									}
								} else {
									if csel.Tag != "" {
										break parseLoop
									}

									csel.Tag = string(blk)
									if psud := strings.Index(csel.Tag, ":"); psud != -1 {
										psuedo := csel.Tag[psud:]
										csel.Tag = csel.Tag[:psud]
										csel.Psuedo = psuedo
									}
								}
							}

							break defaultLoop
						}

						item := items[index]

						switch item {
						case space, coma, hash, dot, bracket, endbracket:
							if doChildren {
								child.Tag = string(blk)
								if psud := strings.Index(child.Tag, ":"); psud != -1 {
									psuedo := child.Tag[psud:]
									child.Tag = child.Tag[:psud]
									child.Psuedo = psuedo
								}
							} else {
								csel.Tag = string(blk)
								if psud := strings.Index(csel.Tag, ":"); psud != -1 {
									psuedo := csel.Tag[psud:]
									csel.Tag = csel.Tag[:psud]
									csel.Psuedo = psuedo
								}
							}

							continue parseLoop
						case orderOpen:

							var order []byte
							order = append(order, item)

							{
							defaultOrderLoop:
								for {
									index++

									item = items[index]
									if index >= itemsLen {
										reodered := string(order)
										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										insertos := strings.Split(ordered, ",")
										if len(insertos) == 1 && insertos[0] == ordered {
											if strings.Contains(string(blk), ":") {
												blk = append(blk, []byte(reodered)...)
											}

											continue defaultLoop
										}

										for _, or := range insertos {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue defaultLoop
									}

									switch item {
									case orderClosed:
										order = append(order, item)

										reordered := string(order)
										ordered := string(order)
										ordered = strings.TrimPrefix(ordered, "(")
										ordered = strings.TrimSuffix(ordered, ")")

										if doChildren {
											if child.Order == nil {
												child.Order = make(map[string]string)
											}
										} else {
											if csel.Order == nil {
												csel.Order = make(map[string]string)
											}
										}

										insertos := strings.Split(ordered, ",")
										if len(insertos) == 1 && insertos[0] == ordered {
											if strings.Contains(string(blk), ":") {
												blk = append(blk, []byte(reordered)...)
											}

											continue defaultLoop
										}

										for _, or := range insertos {
											ors := strings.Split(or, ":")
											if len(ors) < 2 {
												continue
											}

											if doChildren {
												child.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											} else {
												csel.Order[strings.TrimSpace(ors[0])] = strings.TrimSpace(ors[1])
											}
										}

										continue defaultLoop
									default:
										order = append(order, item)
										continue defaultOrderLoop
									}
								}
							}

						}

						blk = append(blk, item)
					}

					continue parseLoop
				}

			}

			index++
		}
	}

	return sels
}

var (
	exactMatch           = "="
	exactWordInListMatch = "~="
	beginOrExactlyMatch  = "|="
	prefixMatch          = "^="
	suffixMatch          = "$="
	containsMatch        = "*="
)

// splitBracketSelector returns the attribute name, value and operator if available
// for attribute selector.
func (queryCtrl) splitBracketSelector(sel string) (string, string, string) {
	sel = strings.TrimPrefix(sel, "[")
	sel = strings.TrimSuffix(sel, "]")

	switch {
	case strings.Contains(sel, exactWordInListMatch):
		splits := strings.Split(sel, exactWordInListMatch)
		if len(splits) == 1 {
			return splits[0], "", exactWordInListMatch
		}

		return splits[0], splits[1], exactWordInListMatch

	case strings.Contains(sel, beginOrExactlyMatch):
		splits := strings.Split(sel, beginOrExactlyMatch)
		if len(splits) == 1 {
			return splits[0], "", beginOrExactlyMatch
		}

		return splits[0], splits[1], beginOrExactlyMatch

	case strings.Contains(sel, prefixMatch):
		splits := strings.Split(sel, prefixMatch)
		if len(splits) == 1 {
			return splits[0], "", prefixMatch
		}

		return splits[0], splits[1], prefixMatch

	case strings.Contains(sel, suffixMatch):
		splits := strings.Split(sel, suffixMatch)
		if len(splits) == 1 {
			return splits[0], "", suffixMatch
		}

		return splits[0], splits[1], suffixMatch

	case strings.Contains(sel, containsMatch):
		splits := strings.Split(sel, containsMatch)
		if len(splits) == 1 {
			return splits[0], "", containsMatch
		}

		return splits[0], splits[1], containsMatch

	case strings.Contains(sel, exactMatch):
		splits := strings.Split(sel, exactMatch)
		if len(splits) == 1 {
			return splits[0], "", exactMatch
		}

		return splits[0], splits[1], exactMatch
	}

	return sel, "", ""
}

func (queryCtrl) attrFor(target *Markup, attrName string, attrVal string, op string) bool {
	attr, err := GetAttr(target, attrName)
	if err != nil {
		return false
	}

	_, val := attr.Render()

	switch op {
	case exactMatch:
		return val == attrVal

	case exactWordInListMatch:
		splits := strings.Split(val, " ")
		for _, item := range splits {
			if item != attrVal {
				continue
			}

			return true
		}

	case beginOrExactlyMatch:
		if strings.HasPrefix(val, attrVal+"-") {
			return true
		}

		return val == attrVal

	case prefixMatch:
		return strings.HasPrefix(val, attrVal)

	case suffixMatch:
		return strings.HasSuffix(val, attrVal)

	case containsMatch:
		return strings.Contains(val, attrVal)

	default:
		return true
	}

	return false
}

func (queryCtrl) tagFor(target *Markup, tag string) bool {
	return target.tagname == tag
}

func (queryCtrl) classFor(target *Markup, class string) bool {
	attr, err := GetAttr(target, "class")
	if err != nil {
		return false
	}

	_, val := attr.Render()
	return strings.Contains(val, class)
}

func (queryCtrl) idFor(target *Markup, id string) bool {
	attr, err := GetAttr(target, "id")
	if err != nil {
		return false
	}

	if _, val := attr.Render(); val == id {
		return true
	}

	return false
}
