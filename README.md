# Gravitee Doc Generation

## Run the example project

Requirements:
* go 1.23

Build the tool
```shell
make build
```

Set up the config
```shell
export DOCGEN_ROOT="$(pwd)/examples/test/config"
```

Run it
```shell
cd examples/test/plugin
../../../bin/doc-gen --write
```

The docs are already generated in `example/plugin` but you can edit README.md above the generation marker.

You can `--dryrun` to print the result in the console or `--validate` to validate run the generation without any outputs.

## Overview

## Core

It has the following features

- Ability to declare more of less required "chunks" of documentation
- Ability to arrange those chunks into one or several output templates
- Ability to specialise the output template for a type of plugin (e.g policy)  or single plugin
- Ability to extend this "core" with extensions
- Ability to keep existing custom content in the target file (limited to the top of the file) 
- Usable as a lib or cli

Configuration consist of two parts

- Files present in the project that must be filled by the writer (dev team, doc team…)
- Chunks and outputs that are predefined and rule files that must be present in the project

In the [repo example directory](https://github.com/gravitee-io-labs/readme-gen) you  will find `plugin` and `config` directory. `plugin` is what you user can/should to and `config` is what defines document structures are outputs.

Chunks are defined with

- target template to layout the data or to be used as is (e.g a Markdown document)
    - path are relative to the plugin repository unless they are absolute
    - most of those file are 
- if it is required
- data type: point of extension to enable processing data injection in templates
- data: configuration for the data type, no particular requirements 

Outputs are define with :

- Master template in which all chunk will laid out
- Target file
- If the target file can keep a part of its existing content  

## Extension (aka data type)

- `table`
    - ability to define columns
    - simple data structure that int consumable in a template
- `options`
    - can generate a table of all configurable attributes from a json schema (with gravitee extension)
- `code`
    - include code in a specific language from one or severals snippets
    - a file can be included before and after those snippets
- `raw-examples`
    - user can define file with plugin custom config exemples
    - title and description can be added
    - json or yaml (CRD) are supported
    - user code is validated against the json schema
    - code is included into a predefined template specialised by plugin types
- `gen-examples`
    - selection of generated example in json or yaml from a schema default values and examples (require to set examples for required value not having defaults)
- `schema-to-yaml` (alpha)
    - turns a json schema into documented Yaml configuration file with comments
        - requires default value or examples for all attributes (not implemented yet)
- `schema-to-env` (alpha)
  - turns a json schema into documented ENV vars configuration file with comments
    - requires default value or examples for all attributes (not implemented yet)

## Other extension

They are easy to implement, two functions: 1 to validate the `data` part of the chuck and one to handle the generation, they then need to be register in `main.go`


### Go template

As part of this tool some go templates have been added to ease formatting

- `default`
    - set a value is the one target is not present
- `ternary`
    - output something for true or false
- `indent` 
    - indent code at the right level
- `pad` 
    - add some padding into code 
- `quote` 
    - add single quote on strings
- `icz` 
    - increase the given value
- `joinset`
    - joins a set (`map[any]bool`) with separator and surround the string values with user input
- `title` 
    - Upper case the first letter
- `mvmdheader`
    -  Include a md file and move their titles level
