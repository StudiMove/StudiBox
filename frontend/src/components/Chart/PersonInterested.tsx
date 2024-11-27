import { useState } from 'react';
import { Bar } from 'react-chartjs-2';
import {
  Chart,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

// Enregistrement des éléments nécessaires pour le graphique
Chart.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const PersonInterested = () => {
  // Jeu de données pour chaque filtre (Tout, Année, Mois, Jour)
  const dataSets = {
    Tout: [1000, 6000, 4000, 8000, 3000, 5000],
    Année: [12000, 14000, 16000, 18000, 20000, 22000],
    Mois: [500, 3000, 4500, 6000, 1000, 8000],
    Jour: [200, 500, 700, 1000, 300, 400],
  };

  const [activeFilter, setActiveFilter] =
    useState<keyof typeof dataSets>('Mois');
  const chartData = dataSets[activeFilter];

  const data = {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
    datasets: [
      {
        label: 'Intérêt (Bleu)',
        data: chartData,
        backgroundColor: '#3498db',
      },
      {
        label: 'Intérêt (Rouge)',
        data: chartData.map((value) => value - 1000), // Juste pour l'exemple
        backgroundColor: '#e74c3c',
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: `Statistiques des personnes intéressées (${activeFilter})`,
      },
    },
  };

  return (
    <div className="bg-white rounded-xl p-6 shadow-md w-full">
      {/* Titre avec icône */}
      <div className="flex items-center mb-4">
        <img src="/path/to/icon.svg" alt="icon" className="h-6 w-6 mr-2" />
        <h3 className="text-xl font-semibold text-black">
          Personne intéressée
        </h3>
      </div>

      {/* Filtres */}
      <div className="flex justify-around mb-4">
        {Object.keys(dataSets).map((filter) => (
          <button
            key={filter}
            onClick={() => setActiveFilter(filter as keyof typeof dataSets)}
            className={`px-4 py-2 text-sm font-semibold ${
              activeFilter === filter
                ? 'text-black border-b-2 border-black'
                : 'text-gray-500'
            }`}
          >
            {filter}
          </button>
        ))}
      </div>

      {/* Graphique à barres */}
      <div className="relative w-full h-64">
        <Bar data={data} options={options} />
      </div>
    </div>
  );
};

export default PersonInterested;
