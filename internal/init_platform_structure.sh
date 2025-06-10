#!/bin/bash

# Define folders
folders=(
  handlers
  models
  repositories
  services
  routes
  utils
)

# Define files for each folder
declare -A files

files[handlers]="auth_handler.go user_handler.go role_handler.go lead_handler.go deal_handler.go task_handler.go document_handler.go sms_handler.go message_handler.go"
files[models]="user.go role.go lead.go deal.go task.go document.go sms_confirmation.go message.go"
files[repositories]="user_repository.go role_repository.go lead_repository.go deal_repository.go task_repository.go document_repository.go sms_confirmation_repository.go message_repository.go"
files[services]="auth_service.go user_service.go role_service.go lead_service.go deal_service.go task_service.go document_service.go sms_service.go message_service.go"

# Create folders and files
for folder in "${folders[@]}"; do
  mkdir -p "$folder"
  echo "📁 Created $folder/"
  for file in ${files[$folder]}; do
    touch "$folder/$file"
    echo "  └── 📄 $file"
  done
done

echo -e "\n✅ Platform structure generated successfully."
