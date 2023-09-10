package docs

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"

	"github.com/delaneyj/gomponents-iconify/iconify/carbon"
	"github.com/delaneyj/gomponents-iconify/iconify/game_icons"
	"github.com/delaneyj/gomponents-iconify/iconify/gridicons"
	"github.com/delaneyj/gomponents-iconify/iconify/material_symbols"
	"github.com/delaneyj/gomponents-iconify/iconify/mdi"
	"github.com/delaneyj/gomponents-iconify/iconify/ph"
	"github.com/delaneyj/gomponents-iconify/iconify/skill_icons"
	"github.com/delaneyj/gomponents-iconify/iconify/tabler"
	"github.com/delaneyj/gomponents-iconify/iconify/vscode_icons"
	"github.com/delaneyj/gomponents-iconify/iconify/zondicons"
	. "github.com/delaneyj/toolbelt/gomps"
	"github.com/dustin/go-humanize"
	"github.com/go-chi/chi/v5"
)

func setupHome(ctx context.Context, router *chi.Mux) error {
	build, err := staticFS.ReadFile("static/datastar.iife.js")
	if err != nil {
		return fmt.Errorf("error reading build: %w", err)
	}
	buf := bytes.NewBuffer(nil)
	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return fmt.Errorf("error compressing build: %w", err)
	}

	if _, err := w.Write(build); err != nil {
		return fmt.Errorf("error compressing build: %w", err)
	}
	w.Close()
	iifeBuildSize := humanize.Bytes(uint64(buf.Len()))

	type Feature struct {
		Description string
		Icon        NODE
		Details     NODE
	}

	features := []Feature{
		{
			Description: "Fine Grained Reactivity via Signals",
			Icon:        ph.GitDiff(),
			Details:     DIV(TXT("No Virtual DOM. proxy wrappers, or re-rendering the entire page on every change.  Take the best available options and use hassle free.")),
		},
		{
			Description: "Fully Compliant",
			Icon:        mdi.LanguageHtmlFive(),
			Details:     DIV(TXT("No monkey patching, no custom elements, no custom attributes, no custom anything.  Just plain old HTML5.")),
		},
		{
			Description: "Everything is an Extension",
			Icon:        gridicons.Plugins(),
			Details:     DIV(TXT("Disagree with the built-in behavior? No problem, just write your own extension in a type safe way.  Take what you need, leave what you don't.")),
		},
		{
			Description: "Batteries Included",
			Icon:        game_icons.Batteries(),
			Details: DIV(
				CLS("breadcrumbs"),
				UL(
					CLS(
						"flex flex-wrap gap-2 justify-center items-center",
					),
					LI(TXT("Actions")),
					LI(TXT("Attribute Binding")),
					LI(TXT("Focus")),
					LI(TXT("HTMX like features")),
					LI(TXT("Intersects")),
					LI(TXT("Two-Way Binding")),
					LI(TXT("Visibility")),
					LI(TXT("Teleporting")),
					LI(TXT("Text Replacement")),
				),
			),
		},
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, HTML5(HTML5Props{
			Title:       "Datastar ",
			Language:    "en",
			Description: `Datastar is a declarative frontend framework that takes the best of modern tooling and combines them with a heavy dose of declarative hypermedia into a single framework that is blazingly easy to use.`,
			Head: NODES{
				LINK(
					REL("icon"),
					HREF(staticPath("favicon.svg")),
				),
				LINK(
					REL("stylesheet"),
					HREF("https://fonts.googleapis.com/css?family=Orbitron|Inter|JetBrains+Mono&display=swap"),
				),
				LINK(
					REL("stylesheet"),
					TYPE("text/css"),
					HREF(staticPath("tailwind.css")),
				),
			},
			Body: NODES{
				CLS("flex flex-col w-full min-h-screen"),
				DIV(
					CLS("p-4 flex flex-col items-center bg-cover bg-opacity-50 text-white bg-center"),
					STYLE(fmt.Sprintf("background-image: url(%s);", staticPath("bg.jpg"))),
					DIV(
						CLS("w-full flex justify-between items-center gap-2  backdrop-blur-sm py-2"),
						DIV(
							CLS("flex gap-2 items-center text-5xl font-display"),
							TXT("Datastar"),
							material_symbols.AwardStarOutline(),
						),
						DIV(
							TXT("Declarative Frontend Framework"),
						),
					),
				),
				DIV(
					CLS("flex justify-end gap-6 px-4 py-1 bg-base-100 text-base-content text-sm"),
					A(
						CLS("btn btn-primary btn-ghost btn-sm"),
						TXT("Docs"),
						HREF("/docs"),
					),
					A(
						CLS("btn btn-primary btn-ghost btn-sm"),
						TXT("Essays"),
						HREF("/essays"),
					),
					DIV(
						CLS("join"),
						A(
							CLS("btn btn-ghost btn-sm"),
							skill_icons.Discord(CLS("text-2xl")),
							HREF("https://discord.com/channels/1035247242887561326/1149367785374359613"),
						),
						A(
							CLS("btn btn-ghost btn-sm"),
							skill_icons.GithubLight(CLS("text-2xl")),
							HREF("https://github.com/delaneyj/datastar"),
						),
					),
				),
				DIV(
					CLS("flex-1 flex flex-wrap md:p-16 text-xl flex-col items-center text-center bg-gradient-to-tr from-base-100 to-base-200"),

					DIV(
						CLS("max-w-4xl flex flex-col items-center justify-center gap-16"),
						DIV(
							CLS("flex flex-wrap gap-4 justify-center items-center text-6xl"),
							vscode_icons.FileTypeAssembly(),
							vscode_icons.FileTypeC(),
							vscode_icons.FileTypeCpp(),
							vscode_icons.FileTypeCobol(),
							vscode_icons.FileTypeClojure(),
							vscode_icons.FileTypeCrystal(),
							vscode_icons.FileTypeCsharp(),
							vscode_icons.FileTypeElixir(),
							vscode_icons.FileTypeFsharp(),
							vscode_icons.FileTypeFortran(),
							vscode_icons.FileTypeGoGopher(),
							vscode_icons.FileTypeHaskell(),
							vscode_icons.FileTypeJava(),
							vscode_icons.FileTypeJs(),
							vscode_icons.FileTypeJulia(),
							vscode_icons.FileTypeKotlin(),
							vscode_icons.FileTypeLisp(),
							vscode_icons.FileTypeLua(),
							vscode_icons.FileTypeNim(),
							vscode_icons.FileTypeOcaml(),
							vscode_icons.FileTypePerl(),
							vscode_icons.FileTypePhp(),
							vscode_icons.FileTypePython(),
							vscode_icons.FileTypeR(),
							vscode_icons.FileTypeRuby(),
							vscode_icons.FileTypeRust(),
							vscode_icons.FileTypeScala(),
							vscode_icons.FileTypeShell(),
							vscode_icons.FileTypeSwift(),
							vscode_icons.FileTypeTypescript(),
							vscode_icons.FileTypeVb(),
							vscode_icons.FileTypeZig(),
						),
						DIV(

							H1(
								CLS("text-6xl font-bold"),
								TXT("HTML on whatever you like"),
							),
							A(
								CLS("link-accent text-4xl"),
								HREF("https://htmx.org/essays/hypermedia-on-whatever-youd-like/"),
								TXT("It's the best idea since web rings"),
							),
						),
						DIV(
							CLS("flex flex-col gap-2 w-full"),
							DIV(
								CLS("bg-base-100 shadow-inner text-base-content p-4 rounded-box"),
								HIGHLIGHT("html", `<div data-signal-count="0">
	<div>
		<button data-on-click="$count++">Increment +</button>
		<button data-on-click="$count--">Decrement -</button>
		<input type="number" data-model="count" />
	</div>
	<button
		data-signal-get="'/api/echo'"
		data-on-load="@get"
		data-on-click="@get"
	>
		Contents
	</button>
</div>
		`,
								),
							),

							DIV(
								CLS("flex gap-2 justify-center items-center"),
								DIV(
									CLS("badge badge-primary flex-1 gap-1"),
									tabler.FileZip(),
									TXT(iifeBuildSize+" w/ all extensions"),
								),
								DIV(
									CLS("badge badge-primary flex-1 gap-1"),
									carbon.ColumnDependency(),
									TXT("0 Dependencies"),
								),
								DIV(
									CLS("badge badge-primary flex-1 gap-1"),
									zondicons.Checkmark(),
									TXT("Fully Tree Shakeable"),
								),
							),
						),
						P(
							TXT("Takes the best of modern tooling and combines them with a heavy dose of declarative hypermedia into a single framework that is blazingly easy to use."),
						),
						DIV(
							CLS("card w-full shadow-2xl ring-4 bg-base-300 ring-secondary text-secondary-content"),
							DIV(
								CLS("card-body flex flex-col justify-center items-center"),
								UL(
									CLS("flex flex-col gap-6 justify-center items-center text-2xl gap-4  max-w-xl"),
									RANGE(features, func(f Feature) NODE {
										return LI(
											DIV(
												CLS("flex flex-col gap-1 justify-center items-center"),
												DIV(
													CLS("flex gap-2 items-center"),
													f.Icon,
													TXT(f.Description),
												),
												DIV(
													CLS("text-lg opacity-50 p-2 rounded"),

													f.Details,
												),
											),
										)
									}),
								),
							),
						),

						DIV(
							CLS("flex flex-col gap-2 justify-center items-center"),
							TXT("Built with "),
							DIV(
								CLS("flex gap-1 justify-center items-center text-5xl"),
								vscode_icons.FileTypeHtml(),
								material_symbols.AddRounded(),
								vscode_icons.FileTypeTypescriptOfficial(),
								material_symbols.AddRounded(),
								vscode_icons.FileTypeVite(),
								material_symbols.AddRounded(),
								vscode_icons.FileTypeGoGopher(),
							),
							DIV(
								CLS("flex gap-2 justify-center items-center"),
								TXT("by "),
								A(
									CLS("link-accent"),
									HREF("http://github.com/delaneyj"),
									TXT("Delaney"),
								),
								TXT("and looking for contributors!"),
							),
						),
						DIV(
							CLS("w-full flex gap-2 items-center"),
							A(
								CLS("btn btn-lg flex-1"),
								HREF("/essays/why-another-framework"),
								material_symbols.Help(),
								TXT("Why another framework?"),
							),
							A(
								CLS("btn btn-primary btn-lg flex-1"),
								HREF("/docs"),
								mdi.RocketLaunch(),
								TXT("Don't care, just get started"),
							),
						),
					),
				),
			},
		}))

	})

	return nil
}
