package main

var (
	// DefaultEnvs is a map of default environment manifests
	DefaultEnvs = map[string]map[string][]string{
		// default is the default environment setup for a basic localhost universe
		"default": map[string][]string{
			// A function is a stateless piece of software, that usually only does
			// one or two things, and has a small footprint
			"functions": []string{
				"sbh 9001 https://gitlab.com/oglinuk/sbh.git",
				"quotitioner 9429 https://gitlab.com/oglinuk/quotitioner.git",
			},

			// An application is a piece of software that requires state, usually
			// having a medium-large footprint
			"applications": []string{
				// TODO: Fix (to 9999) when gendockercompose is implemented
				"gitea 3000 https://github.com/go-gitea/gitea.git",
			},

			// A blog is a git hosted collection of code/.md/.html/site/... files around
			// discussion of a topic
			// TODO: Split blogs into individualsBlog/projectsBlog?
			"blogs": []string{
				"fragglet-blog 50006 https://github.com/fragglet/soulsphere.org.git",
				"rwxrob-blog 50007 https://github.com/rwxrob/rwxrob.git",
				"jessfraz-blog 50008 https://github.com/jessfraz/blog.git",
				"gitea-blog 50300 https://gitea.com/gitea/blog.git",
			},

			// A doc source is a git hosted collection of code/.md/.html/.txt/... files
			// documenting or explaining something
			"docs": []string{
				"rwx.gg 50500 https://gitlab.com/rwx.gg/README.git",
				"freedoom-docs 50501 https://github.com/freedoom/freedoom.github.io.git",
			},

			// A fileserver source is a git hosted repo that does not necessarily fall
			// into the docs or blogs categories, but still requires the same scaffolding
			"fileserver": []string{
				"mediafs 12121 https://gitlab.com/oglinuk/mediafs.git",
				"directories 50000 https://gitlab.com/oglinuk/directories.git",
				"library 50001 https://gitlab.com/oglinuk/library.git",
				"ptp 50002 https://github.com/oglinuk/ptp.git",
				"awesome-gitea 50700 https://gitea.com/gitea/awesome-gitea.git",
			},

			// A game source is a git hosted game which is capable of being run locally
			// TODO: Dockerize?
			"games": []string{
				"chocolate-doom 30303 https://github.com/chocolate-doom/chocolate-doom.git",
				"quakespasm 30304 https://github.com/ericwa/quakespasm.git",
			},

			// A personal source is ones creation(s) that falls under any type of service
			"personal": []string{
				//"fourohfournotfound 40404 https://gitlab.com/oglinuk/fourohfournotfound.git"
			},
		},
	}
)
