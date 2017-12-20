package mgo

import (
	"fmt"
	goast "go/ast"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/mgokit/static"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// MongoGen generates a mongodb based CRUD api for a struct declaration.
func MongoGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	var hasPublicID bool

	// Validate we have a `PublicID` field.
	{
	fieldLoop:
		for _, field := range str.Struct.Fields.List {
			typeIdent, ok := field.Type.(*goast.Ident)
			if !ok {
				continue
			}

			// If typeName is not a string, skip.
			if typeIdent.Name != "string" {
				continue
			}

			for _, indent := range field.Names {
				if indent.Name == "PublicID" {
					hasPublicID = true
					break fieldLoop
				}
			}
		}
	}

	if !hasPublicID {
		return nil, fmt.Errorf(`Struct has no 'PublicID' field with 'string' type
		 Add 'PublicID string json:"public_id"' to struct %q
		`, str.Object.Name.Name)
	}

	packageName := fmt.Sprintf("%smgo", strings.ToLower(str.Object.Name.Name))
	packageFinalPath := filepath.Join(toPackage, packageName)

	//mongoTestGen := gen.Block(
	//	gen.Package(
	//		gen.Name(fmt.Sprintf("%s_test", packageName)),
	//		gen.Imports(
	//			gen.Import("os", ""),
	//			gen.Import("time", ""),
	//			gen.Import("testing", ""),
	//			gen.Import("gopkg.in/mgo.v2", "mgo"),
	//			gen.Import("github.com/influx6/faux/tests", ""),
	//			gen.Import("github.com/influx6/faux/metrics", ""),
	//			gen.Import("github.com/influx6/faux/context", ""),
	//			gen.Import("github.com/influx6/faux/db/mongo", ""),
	//			gen.Import("github.com/influx6/faux/metrics/custom", ""),
	//			gen.Import(packageFinalPath, "mdb"),
	//			gen.Import(str.Path, ""),
	//		),
	//		gen.Block(
	//			gen.SourceTextWith(
	//				string(static.MustReadFile("mongo-api-test.tml", true)),
	//				gen.ToTemplateFuncs(
	//					ast.ASTTemplatFuncs,
	//					template.FuncMap{
	//						"map":       ast.MapOutFields,
	//						"mapValues": ast.MapOutValues,
	//						"hasFunc":   pkgDeclr.HasFunctionFor,
	//					},
	//				),
	//				struct {
	//					Pkg          *ast.PackageDeclaration
	//					Struct       ast.StructDeclaration
	//				}{
	//					Pkg:          &pkgDeclr,
	//					Struct:       str,
	//				},
	//			),
	//		),
	//	),
	//)

	mongoReadmeGen := gen.Block(
		gen.Block(
			gen.SourceText(
				string(static.MustReadFile("mongo-api-readme.tml", true)),
				struct {
					Pkg          *ast.PackageDeclaration
					Struct       ast.StructDeclaration
					CreateAction ast.StructDeclaration
					UpdateAction ast.StructDeclaration
					PackageName  string
					PackagePath  string
				}{
					PackagePath: packageFinalPath,
					PackageName: packageName,
					Pkg:         &pkgDeclr,
					Struct:      str,
				},
			),
		),
	)

	mongoJSONGen := gen.Block(
		gen.Package(
			gen.Name("fixtures"),
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
							"map":           ast.MapOutFields,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
							"mapValues":     ast.MapOutValues,
							"mapJSON":       ast.MapOutFieldsToJSON,
							"hasFunc":       pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
					},
				),
			),
		),
	)

	mongoBackendGen := gen.Block(
		gen.Package(
			gen.Name("types"),
			gen.Imports(
				gen.Import("github.com/influx6/faux/context", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("mongo-api-backend.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"map":     ast.MapOutFields,
							"hasFunc": pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
					},
				),
			),
		),
	)

	mongoGen := gen.Block(
		gen.Commentary(
			gen.SourceText(`Package `+packageName+` provides a auto-generated package which contains a sql CRUD API for the specific {{.Object.Name}} struct in package {{.Package}}.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("strings", ""),
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
							"hasFunc": pkgDeclr.HasFunctionFor,
						},
					),
					struct {
						Pkg    *ast.PackageDeclaration
						Struct ast.StructDeclaration
					}{
						Pkg:    &pkgDeclr,
						Struct: str,
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
		},
		{
			Writer:       fmtwriter.New(mongoBackendGen, true, true),
			FileName:     fmt.Sprintf("%s_backend.go", strings.ToLower(str.Object.Name.Name)),
			Dir:          "types",
			DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(mongoGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
		{
			Writer:       mongoJSONGen,
			FileName:     fmt.Sprintf("%s_fixtures.go", packageName),
			Dir:          filepath.Join(packageName, "fixtures"),
			DontOverride: true,
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
						"hasFunc": pkgDeclr.HasFunctionFor,
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
