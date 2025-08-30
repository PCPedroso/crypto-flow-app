# INFRASTRUCTURE — IaC e comandos

## Onde está o IaC
- Pasta exemplo: /infra (Terraform / ARM / Bicep / CloudFormation)

## Comandos comuns
- Terraform init/plan/apply:
  terraform init
  terraform plan -out plan.tfplan
  terraform apply plan.tfplan

## Infra local
- docker-compose up --build
- Scripts para levantar serviços de dev (ver /scripts)

## Notas
- Documentar state backend (S3/GCS) e locks