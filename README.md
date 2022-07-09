# tw5-json-markdown

Converts the single JSON file exported from [TiddlyWiki](https://tiddlywiki.com) version 5 to individual Markdown files.


## Use case

If you have a TiddlyWiki containing TiddlyWiki 5 tiddlers and want to export them to Markdown, perhaps to convert to a website with the [Hugo](https://gohugo.io/) static site generator, there is no built-in export. However TiddlyWiki does allow all of the tiddlers to be exported into a single JSON file containing the metadata, tags etc.

tw5-json-markdown takes this JSON file and generates a folder of Markdown-format files from the tiddler data in it.


## Exporting to JSON

* Open the **Advanced Search** tiddler
* Click on the **Filters** tab
* Enter `[!is[system]]` to list all of the non-system tiddlers
* Click on the **Export** button to export every tiddler to a single JSON file


## Converting the JSON file to Markdown

The executable takes two parameters: `tw5-json-markdown -in <jsonFile> -out <markdownFolder>`

A level one header containing the tiddler title is added to the top as the tiddler title is rendered in TiddlyWiki from metadata.

## Supported formatting

* Numbered lists
* Bulleted lists
* Font formats: bold, italic, raw, underline, strikeout, superscript, subscript
* Headings
* Blockquotes using `> format`
* Links in `[[double bracket]]` format including `[[shown|link]]` and `[ext[shown|https://example.com]]`
* Transclusion (requires MultiMarkdown to render the generated Markdown)
* Tables

## Limitations

* Table alignment isn't supported yet
* Automatic CamelCase WikiLinks are not supported as they appear to be deprecated in TiddlyWiki

## Planned support

* [ ] Table alignment
