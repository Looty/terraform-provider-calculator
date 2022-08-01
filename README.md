# terraform-provider-calculator
Terraform provider for Go Calculator

## How to run
- Inside `root` directory:
    - run `make install`
- Inside `examples/` directory, run:
  ```bash
    make clean
    tf init
    tf apply --auto-approve
    ```

## Usage
```terraform
terraform {
  required_providers {
    calculator = {
      version = "0.0.1"
      source  = "hashicorp.com/Looty/calculator"
    }
  }
}

provider "calculator" {}

data "calculator_compute" "example" {
  a = 3
  b = 6

  function = "div"
}

output "example_result" {
  value = data.calculator_compute.example.result
}
```