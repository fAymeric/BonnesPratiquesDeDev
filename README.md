# 🔤 HangmanWeb - Jeu du Pendu en Go

---

## 📝 Description

**HangmanWeb** est un jeu web simple du Pendu, développé en Go côté serveur. Il permet de jouer directement dans un navigateur avec un rendu dynamique et une logique de jeu totalement gérée côté backend.

---

## 🌟 Fonctionnalités

- 🎮 Jeu du pendu interactif dans le navigateur
- 🧠 Sélection aléatoire de mots depuis une liste
- 🧾 Affichage dynamique du mot à deviner, des lettres proposées
- ✅ Gestion des victoires/défaites avec messages finaux
- 🗃️ Score board (optionnel selon implémentation)
- 🌐 Choix de langues et de difficultés 

---

## 🛠️ Stack Technique

- **Langage backend** : [Go (Golang)](https://golang.org/) 🦫
- **Serveur HTTP** : `net/http` intégré à Go
- **Frontend** : HTML, CSS & JavaScript vanilla
- **Templates** : HTML templates (`html/template`)

---

## ⚙️ Installation & Lancement


### **Étapes d’installation**

1. **Cloner le dépôt :**

   ```bash
   git clone https://github.com/fAymeric/BonnesPratiquesDeDev.git
   cd BonnesPratiquesDeDev/hangman_web
   go run ./cmd/main.go
