# Main.go

* For mapping file
  * parse + aggregate into structured list of mapping data
* Open `schema.json`, parse into structured data
* For each mapping file
  * For each mapping in the file
     * autogen the a ruleset for the mapping
     * autogen the docs for that mapping

* Autogen `rules/apispec/provider.go`
* Autogen `docs/README.md`


# Missing
* Unit Test this code
* Document this program
* Periodically genrate rule set
  * [this repo] Generate periodic releases using nightly trigger
  * [tf lint] Check for latest release + warn if there are updates, possible to use the SHA hash of the binary

* Pin provider to API spec



## Proposal for pinning to TF Provider Version

> Assuming `PROVIDER_VERSION` is the `azurerm` tf provider version

* `tools/apispec-rule-gen/mappings` could become `tools/apispec-rule-gen/mappings-$PROVIDER_VERSION`
* `tools/apispec-rule-gen/schema` could become `tools/apispec-rule-gen/schema-$PROVIDER_VERSION`
* `tools/apispec-rule-gen/main.go` could iterate over:
  * all `$PROVIDER_VERSION` mappings + schemas

* Manually created rules could live in `rules-$PROVIDER_VERSION/`
* The tool will need to generate
  * `docs-$PROVIDER_VERSION/rules`
  * `rules-$PROVIDER_VERSION/apispec`

* Releases could look like