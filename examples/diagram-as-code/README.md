# example

<div style="text-align:center"><img src="example.drawio.svg" /></div>

- [diagram.xml](diagram.xml)
- [structure.config.yaml](structure.config.yaml)

## Commands

User guide: 

```bash
$ gcp-terraform-generator --workdir .
```

Commands:

```bash
$ gcp-terraform-generator diagram -c diagram.config.yaml -d diagram.xml -o diagram.yaml
$ gcp-terraform-generator structure -c structure.config.yaml -o ./output
$ gcp-terraform-generator code -c diagram.yaml -o ./output
```