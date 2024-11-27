import { Chart } from 'react-chartjs-2';
import { useState } from 'react';

const SalesStatsComponent = () => {
  const [activeTab, setActiveTab] = useState('Billets'); // Par défaut, l'onglet 'Billets'

  const handleTabClick = (tab: string) => {
    setActiveTab(tab);
  };

  // Le jeu de données pour le graphique mixte
  const data = {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'], // Labels sur l'axe des X
    datasets: [
      {
        type: 'bar' as const, // Type de graphique (barre)
        label: 'Recette totale (barre)', // Label du jeu de données
        data: [50, 10000, 7500, 15000, 20000, 25000], // Données des barres
        backgroundColor: 'rgba(54, 162, 235, 0.6)', // Couleur de remplissage des barres
        borderColor: 'rgba(54, 162, 235, 1)', // Couleur des bordures
        borderWidth: 1, // Largeur des bordures
      },
      {
        type: 'line' as const, // Type de graphique (ligne)
        label: 'Nombre de vues (ligne)', // Label du jeu de données pour la ligne
        data: [10000, 15000, 10000, 17000, 21000, 24000], // Données de la ligne
        fill: false, // Pas de remplissage sous la ligne
        borderColor: '#e74c3c', // Couleur de la ligne
        tension: 0.1, // Tension pour arrondir les angles
        borderWidth: 2, // Largeur de la ligne
      },
    ],
  };

  // Options pour le graphique
  const options = {
    responsive: true, // Rendre le graphique responsive
    maintainAspectRatio: false, // Permet de mieux gérer la taille du graphique
    plugins: {
      legend: {
        position: 'bottom' as const, // Position de la légende
      },
    },
    scales: {
      x: {
        beginAtZero: true, // Assure que l'axe des X commence à zéro
      },
      y: {
        beginAtZero: true, // Assure que l'axe des Y commence à zéro
        max: 25000, // Définit une valeur max pour éviter que la ligne sorte du graphique
      },
    },
  };

  return (
    <div className="bg-white rounded-xl p-6 shadow-md w-full">
      {/* Onglets */}
      <div className="flex space-x-8 mb-4">
        {['Billets', 'Options', 'Recette', 'Clic vers lien'].map((tab) => (
          <button
            key={tab}
            onClick={() => handleTabClick(tab)}
            className={`px-4 py-2 font-semibold ${
              activeTab === tab
                ? 'text-black border-b-2 border-black'
                : 'text-gray-500'
            }`}
          >
            {tab}
          </button>
        ))}
      </div>

      {/* Titre et navigation */}
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-xl font-semibold text-black">{activeTab}</h3>
        <div className="flex items-center space-x-4">
          <button className="text-sm font-semibold text-gray-500">&lt;</button>
          <span className="text-sm font-semibold">May</span>
          <button className="text-sm font-semibold text-gray-500">&gt;</button>
        </div>
      </div>

      {/* Graphique mixte */}
      <div className="relative w-full h-64">
        {/* Ajout du type global comme 'bar' */}
        <Chart type="bar" data={data} options={options} />
      </div>
    </div>
  );
};

export default SalesStatsComponent;
