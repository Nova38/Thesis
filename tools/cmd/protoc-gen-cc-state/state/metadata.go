package state

import (
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

type CCMetadataModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func CCModule() *CCMetadataModule { return &CCMetadataModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *CCMetadataModule) Name() string { return "cc-metadata" }

func (p *CCMetadataModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())
}

//func (p *CCMetadataModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
//	for _, f := range targets {
//		if f.BuildTarget() {
//			p.generateFile(f)
//		}
//	}
//	return p.Artifacts()
//}
