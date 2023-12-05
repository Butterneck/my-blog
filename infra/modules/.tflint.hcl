plugin "terraform" {
  enabled = true
  // preset  = "recommended"
  preset  = "all"
}

plugin "aws" {
    enabled = true
    version = "0.28.0"
    source  = "github.com/terraform-linters/tflint-ruleset-aws"
}