package mgo

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/mgokit/static"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// MongoGen generates a mongodb based CRUD api for a struct declaration.
func MongoGen(toDir string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	updateAction := str
	createAction := str

	if len(str.Associations) != 0 {
		if newAction, ok := str.Associations["New"]; ok {
			if action, err := ast.FindStructType(pkgDeclr, newAction.TypeName); err == nil && action.Declr != nil {
				createAction = action
			}
		}

		if upAction, ok := str.Associations["Update"]; ok {
			if action, err := ast.FindStructType(pkgDeclr, upAction.TypeName); err == nil && action.Declr != nil {
				updateAction = action
			}
		}
	}

	packageName := fmt.Sprintf("%sdb", strings.ToLower(str.Object.Name.Name))

	mongoTestGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("os", ""),
				gen.Import("testing", ""),
				gen.Import("encoding/json", ""),
				gen.Import("gopkg.in/mgo.v2", "mgo"),
				gen.Import("gopkg.in/mgo.v2/bson", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/db/mongo", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(filepath.Join(str.Path, toDir, packageName), "mdb"),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("mongo-api-test.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":       ast.MapOutFields,
							"mapValues": ast.MapOutValues,
							"hasFunc":   ast.HasFunctionFor(pkgDeclr),
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	mongoReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText(
				string(static.MustReadFile("mongo-api-readme.tml", true)),
				struct {
					Pkg          *ast.PackageDeclaration
					Struct       ast.StructDeclaration
					CreateAction ast.StructDeclaration
					UpdateAction ast.StructDeclaration
				}{
					Pkg:          &pkgDeclr,
					Struct:       str,
					CreateAction: createAction,
					UpdateAction: updateAction,
				},
			),
		),
	)

	mongoJSONGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("mongo-api-json.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":       ast.MapOutFields,
							"mapValues": ast.MapOutValues,
							"mapJSON":   ast.MapOutFieldsToJSON,
							"hasFunc":   ast.HasFunctionFor(pkgDeclr),
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	mongoGen := gen.Block(
		gen.Commentary(
			gen.SourceText(`Package mdb provides a auto-generated package which contains a mongo CRUD API for the specific {{.Object.Name}} struct in package {{.Package}}.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import("gopkg.in/mgo.v2", "mgo"),
				gen.Import("gopkg.in/mgo.v2/bson", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("mongo-api.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":     ast.MapOutFields,
							"hasFunc": ast.HasFunctionFor(pkgDeclr),
						},
					),
					struct {
						Pkg          *ast.PackageDeclaration
						Struct       ast.StructDeclaration
						CreateAction ast.StructDeclaration
						UpdateAction ast.StructDeclaration
					}{
						Pkg:          &pkgDeclr,
						Struct:       str,
						CreateAction: createAction,
						UpdateAction: updateAction,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:   mongoReadmeGen,
			FileName: "README.md",
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(mongoTestGen, true, true),
			FileName: fmt.Sprintf("%s_test.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(mongoGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
		{
			Writer:   mongoJSONGen,
			FileName: fmt.Sprintf("%s_fixtures_test.go", packageName),
			Dir:      packageName,
			// DontOverride: true,
		},
	}, nil
}

// MongoSolo generates a simple mongo implementation for executing code on mongodb.
func MongoSolo(toDir string, an ast.AnnotationDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	mongoReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText(
				string(static.MustReadFile("mongo-solo-readme.tml", true)),
				struct {
					Pkg     *ast.PackageDeclaration
					Package ast.Package
				}{
					Pkg:     &pkgDeclr,
					Package: pkg,
				},
			),
		),
	)

	mongoGen := gen.Block(
		gen.Commentary(
			gen.Text(`Package mongoapi provides a auto-generated package which contains a mongo base pkg for db operations.`),
		),
		gen.Package(
			gen.Name("mdb"),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import("gopkg.in/mgo.v2", "mgo"),
				gen.Import("gopkg.in/mgo.v2/bson", ""),
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("mongo-solo.tml", true)),
					template.FuncMap{
						"map":     ast.MapOutFields,
						"hasFunc": ast.HasFunctionFor(pkgDeclr),
					},
					struct {
						Pkg     *ast.PackageDeclaration
						Package ast.Package
					}{
						Pkg:     &pkgDeclr,
						Package: pkg,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:   mongoReadmeGen,
			FileName: "README.md",
			Dir:      "mdb",
			// DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(mongoGen, true, true),
			FileName: "mdb.go",
			Dir:      "mdb",
			// DontOverride: true,
		},
	}, nil
}
