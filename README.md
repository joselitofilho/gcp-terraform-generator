<link rel="stylesheet" href="markdown-styles-list.css">

<div align="center">

# GCP Terraform Generator

[![GitHub tag](https://img.shields.io/github/release/joselitofilho/gcp-terraform-generator?include_prereleases=&sort=semver&color=2ea44f&style=for-the-badge)](https://github.com/joselitofilho/gcp-terraform-generator/releases/)
[![Go Report Card](https://goreportcard.com/badge/github.com/joselitofilho/gcp-terraform-generator?style=for-the-badge)](https://goreportcard.com/report/github.com/joselitofilho/gcp-terraform-generator)
[![Code coverage](https://img.shields.io/badge/Coverage-68.2%25-yellow?style=for-the-badge)](#)

[![Made with Golang](https://img.shields.io/badge/Golang-1.21.6-blue?logo=go&logoColor=white&style=for-the-badge)](https://go.dev "Go to Golang homepage")
[![Using Terraform](https://img.shields.io/badge/Terraform-4.84.0-blueviolet?logo=terraform&logoColor=white&style=for-the-badge)](https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs "Go to Terraform docs")
[![Using Diagrams](https://img.shields.io/badge/diagrams.net-orange?logo=&logoColor=white&style=for-the-badge)](https://app.diagrams.net/ "Go to Diagrams homepage")

[![BuyMeACoffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-ffdd00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/joselitofilho)

</div>

# Overview

The GCP Terraform Generator is a powerful tool designed to simplify and streamline the process of creating Terraform configurations for GCP infrastructure. With this tool, you can quickly generate Terraform code to provision GCP resources such as IoT Core, BQ tables, App engines, Cloud functions, Pub Sub, and much more.

[![Start Here](https://img.shields.io/badge/start%20here-blue?style=for-the-badge)](#recommended-step-by-step)

**Table of contents**

- [Install](#install)
- [Third Party Tools](#third-party-tools)
- [Features](#features)
  - [Code generator](#code-generator)
  - [Diagram from code](#diagram-from-code)
  - [Compare diagrams](#compare-diagrams)
- [How it works](#how-it-works)
- [Recommended step by step](#recommended-step-by-step)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

 ```bash
 $ go install github.com/joselitofilho/gcp-terraform-generator@latest
 ```

## Third Party Tools

- [**graphviz**][graphviz]: Graphviz is open source graph visualization software. Graph visualization is a way of representing structural information as diagrams of abstract graphs and networks.
- [**terraform**][terraform]: Terraform is an infrastructure as code tool that lets you build, change, and version cloud and on-prem resources safely and efficiently.

## Features:
- Generate initial stack infrastructure folders.
- Generate GoLang code and Terraform files.
- [*Diagrams*][diagrams] integration: Generate everything based on the exported XML diagram.
- Customization Options: Tailor generated code to your specific requirements using customizable templates and configuration parameters.
- Best Practices: Adhere to GCP and Terraform best practices with automatically generated code that follows industry standards.
- [Supported resources][supported-resources]:
  - [x] App Engine
  - [x] Big Query
  - [x] Big Table
  - [x] DataFlow
  - [x] IoT Core
  - [x] Pub Sub
  - [x] Storage
- Everything is customizable.

### Code generator

<div align="center">

![](assets/code-generator.gif)

</div>

```bash
$ gcp-terraform-generator --workdir ./examples/diagram-as-code
```

## How it works

The code generator already comes with some pre-configured templates for generating Terraform and GoLang files. All generator 
configuration is based on YAML files, making it easy to customize the available resources and templates.

The first step is to write the configuration file specified [here](CONFIGURATION.md). You can also use this [example](examples/diagram-as-code) as a reference.

There you go! Now you can generate the structure of your project or the files based on the configured resources. You can execute the commands in any order.

If you're using [*Diagrams*][diagrams], you can also generate the initial configuration file based on the XML generated by the tool.

If you have any questions or suggestions, go to the [Contributing](#contributing) section. Your contribution is much appreciated.

<div style="text-align:center"><img src="assets/general-overview.svg" /></div>

## Recommended step by step

**Step 1**: Create a folder to organize the diagram and configuration files, ideally named after your stack.
```bash
$ mkdir mystack
```

**Step 2**: Create your diagram using [*Diagrams*][diagrams]. If you have already created one, proceed to the next step.

**Step 3**: Export and download your diagram as an XML file (file name suggestion: `diagram.xml`).
You can find instructions on how to do that at this link: https://www.drawio.com/doc/faq/export-to-xml.

Move the file to the folder created in the Step 1.

```bash
$ mv ~/Downloads/diagram.xml mystack/diagram.xml
```

**Step 4**: Create the diagram configuration file.

Suggestion [diagram.config.yaml](./examples/diagram-as-code/diagram.config.yaml):
```bash
$ cp ./examples/diagram-as-code/diagram.config.yaml mystack/diagram.config.yaml
```

Change the values according to your project.

**Step 5**: Create the structure configuration file.

Suggestion [structure.config.yaml](./examples/diagram-as-code/structure.config.yaml):
```bash
$ cp ./examples/diagram-as-code/structure.config.yaml mystack/structure.config.yaml
```

**Step 6**: Run the generator guide to assist you.

```bash
$ gcp-terraform-generator --workdir mystack
```

Output:
```bash


                 ██████╗ ██████╗ ██████╗ ███████╗     ██████╗ ███████╗███╗   ██╗
                ██╔════╝██╔═══██╗██╔══██╗██╔════╝    ██╔════╝ ██╔════╝████╗  ██║
                ██║     ██║   ██║██║  ██║█████╗      ██║  ███╗█████╗  ██╔██╗ ██║
                ██║     ██║   ██║██║  ██║██╔══╝      ██║   ██║██╔══╝  ██║╚██╗██║
                ╚██████╗╚██████╔╝██████╔╝███████╗    ╚██████╔╝███████╗██║ ╚████║
                 ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝     ╚═════╝ ╚══════╝╚═╝  ╚═══╝
                                                                             GCP


? What would you like to do?  [Use arrows to move, type to filter]
> Generate a diagram config file
  Generate the initial structure
  Generate code
  Exit
```

## Usage

To use these configurations:

1. Navigate to the desired stack/environment folder.
2. Customize the Terraform files (`main.tf`, `vars.tf`, etc.) according to your requirements.
3. Run commands to manage the infrastructure.

User guide: 

```bash
$ gcp-terraform-generator --workdir ./example
```

Or use commands:

```bash
$ gcp-terraform-generator diagram -c ./example/diagram.config.yaml -d ./example/diagram.xml -o ./example/diagram.yaml
$ gcp-terraform-generator structure -c ./example/structure.config.yaml -o ./output
$ gcp-terraform-generator apigateway -c ./example/diagram.yaml -o ./output
$ gcp-terraform-generator lambda -c ./example/diagram.yaml -o ./output/mystack
$ gcp-terraform-generator kinesis -c ./example/diagram.yaml -o ./output/mystack
$ gcp-terraform-generator sqs -c ./example/diagram.yaml -o ./output/mystack
$ gcp-terraform-generator s3 -c ./example/diagram.yaml -o ./output/mystack
```

## Configuration

All you need know regarding configuration you can find in the [configuration](CONFIGURATION.md) section.

[![open - Configuration](https://img.shields.io/badge/open-configuration-blue?style=for-the-badge)](CONFIGURATION.md "Go to configuration")

## Template

For code generation, we are using the standard Golang library [text/template][lib-template]. Further details about the available variables and the definition of some added utility functions can be found in the [template](TEMPLATE.md) section.

[![open - Template](https://img.shields.io/badge/open-template-blue?style=for-the-badge)](TEMPLATE.md "Go to configuration")

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create an [issue][issues] or submit a pull request. Your contribution is much appreciated. See [Contributing](CONTRIBUTING.md).

[![open - Contributing](https://img.shields.io/badge/open-contributing-blue?style=for-the-badge)](CONTRIBUTING.md "Go to contributing")

## License

This project is licensed under the [MIT License](LICENSE).

[diagrams]: https://app.diagrams.net/
[issues]: https://github.com/joselitofilho/gcp-terraform-generator/issues
[graphviz]: https://graphviz.org/download/
[lib-template]: https://pkg.go.dev/text/template
[supported-resources]: https://drive.google.com/file/d/1Lrh6SikW1bvGXrfJLRDFBB4BChQdAPqz/view?usp=sharing
[terraform]: https://developer.hashicorp.com/terraform/tutorials/gcp-get-started/install-cli