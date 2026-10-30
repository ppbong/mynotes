# Github SSH

## 1. Generate SSH Key

```bash
ssh-keygen -t ed25519 -C "name@example.com"
cat ~/.ssh/id_ed25519.pub
```

## 2. Setup SSH Key

- `Github login`
  - `Settings`
    - `SSH and GPG keys`
      - `New SSH key`

## 3. Git with SSH

- new repository

```bash
git init
echo "# example" >> README.md
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:account/repo.git
git push -u origin main
``` 

- old repository
 
```bash
git remote get-url origin
git remote set-url origin git@github.com:account/repo.git
git push
```
