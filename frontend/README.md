# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
export default tseslint.config({
  languageOptions: {
    // other options...
    parserOptions: {
      project: ['./tsconfig.node.json', './tsconfig.app.json'],
      tsconfigRootDir: import.meta.dirname,
    },
  },
});
```

- Replace `tseslint.configs.recommended` to `tseslint.configs.recommendedTypeChecked` or `tseslint.configs.strictTypeChecked`
- Optionally add `...tseslint.configs.stylisticTypeChecked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and update the config:

```js
// eslint.config.js
import react from 'eslint-plugin-react';

export default tseslint.config({
  // Set the react version
  settings: { react: { version: '18.3' } },
  plugins: {
    // Add the react plugin
    react,
  },
  rules: {
    // other rules...
    // Enable its recommended rules
    ...react.configs.recommended.rules,
    ...react.configs['jsx-runtime'].rules,
  },
});
```

frontend/
│
├── public/ # Contient les fichiers statiques comme index.html
│ │ ├── icon/ # Contiendra les fichiers .ico ou autres icônes
│ │ │ ├── fav-icon
├── src/
│ ├── assets/ # Images, fonts, etc.
│ │ ├── branding/ # Contiendra le logo (PNG, SVG, etc.)
│ │ ├── illustration/ # Contiendra des illustrations
│ ├── components/ # Composants réutilisables
│ ├── config/ # Fichiers de configuration spécifiques
│ ├── contexts/ # Contexts pour la gestion de l'état global
│ ├── hooks/ # Hooks personnalisés
│ ├── models/ # Types et interfaces TypeScript
│ └── InputField/
│ └── InputFieldModel.ts
│ ├── pages/ # Pages ou vues principales de l'application
│ ├── services/ # Services pour les appels API et autres
│ ├── states/ # Gestion d'état (Redux, Zustand, etc.)
│ ├── styles/ # Fichiers de styles globaux (CSS, Tailwind, etc.)
│ ├── ├── tailwind.css
│ ├── ├── App.css
│ ├── ├── index.css
│ ├── tests/ # Tests unitaires et d'intégration
│ ├── types/ # Tests unitaires et d'intégration
│ ├── └── InputField/
│ │ │ │ ├── InputFieldProps.d.ts
│ │ │ │ └── InputFieldType.d.ts
│ ├── utils/ # Fonctions utilitaires et helpers
│ ├── App.tsx # Composant principal de l'application
│ ├── main.tsx # Point d'entrée de l'application
│ └── routes.tsx # Configuration des routes
│ └── vite-env.d.ts # Configuration des routes
│
├── .dockerignore # Fichiers et dossiers à ignorer pour Docker
├── .env # Variables d'environnement pour le développement
├── .env.production # Variables d'environnement pour la production
├── .gitignore # Fichiers et dossiers à ignorer pour Git
├── docker-compose.yml # Configuration Docker Compose
├── Dockerfile # Fichier Docker pour construire l'image de l'application
├── eslint.config.js
├──index.html
├── package.json # Dépendances et scripts de l'application
├── postcss.config.js
├── README.md
├── tailwind.config.js
├── tsconfig.app.json
├── tsconfig.json # Configuration TypeScript
├── tsconfig.node.json
└── tsconfig.json # Configuration TypeScript
└── vite.config.ts
