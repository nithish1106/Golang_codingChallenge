# Nithish_Challenge

## Infrastructure Setup

1. Install [Terraform](https://www.terraform.io/), [Ansible](https://www.ansible.com/).
2. Navigate to the `infrastructure/terraform/` folder.
3. Run `terraform init` to initialize.
4. Run `terraform apply` to provision the EC2 instance and security group.
5. Navigate to `infrastructure/ansible/` and run `ansible-playbook playbook.yml` to configure the webserver.

## Coding Challenge

To validate credit card numbers, navigate to the `app/` folder and run:

```bash
go run validate_credit_card.go
