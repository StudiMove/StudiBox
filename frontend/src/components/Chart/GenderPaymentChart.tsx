import { useState } from 'react';
import { Doughnut } from 'react-chartjs-2';
import { Chart, ArcElement, Tooltip, Legend } from 'chart.js';

// Enregistrement des éléments de Chart.js
Chart.register(ArcElement, Tooltip, Legend);

const GenderPaymentChart = () => {
  // Données pour le graphique Civilité
  const genderData = {
    labels: ['Femme', 'Homme', 'Autre'],
    datasets: [
      {
        data: [50, 30, 20], // Exemple de données
        backgroundColor: ['#3498db', '#e74c3c', '#1abc9c'],
        hoverBackgroundColor: ['#2980b9', '#c0392b', '#16a085'],
        borderWidth: 0,
      },
    ],
  };

  // Données pour le graphique Type de paiement
  const paymentData = {
    labels: ['Carte', 'Espèces', 'Chèque'],
    datasets: [
      {
        data: [60, 25, 15], // Exemple de données
        backgroundColor: ['#3498db', '#f1c40f', '#e67e22'],
        hoverBackgroundColor: ['#2980b9', '#f39c12', '#d35400'],
        borderWidth: 0,
      },
    ],
  };

  // Gestion de l'onglet actif (0 = Civilité, 1 = Type de paiement)
  const [activeTab, setActiveTab] = useState(0);

  const options = {
    cutout: '70%', // Crée l'effet d'anneau
    responsive: true,
    plugins: {
      legend: {
        display: false, // On masque la légende pour le graphique
      },
    },
  };

  return (
    <div className="bg-white rounded-xl p-6 shadow-md w-full max-w-md ">
      {/* Onglets */}
      <div className="flex justify-center mb-4">
        <button
          onClick={() => setActiveTab(0)}
          className={`px-4 py-2 text-sm font-semibold ${
            activeTab === 0
              ? 'text-black border-b-2 border-black'
              : 'text-gray-500'
          }`}
        >
          Civilité
        </button>
        <button
          onClick={() => setActiveTab(1)}
          className={`px-4 py-2 text-sm font-semibold ${
            activeTab === 1
              ? 'text-black border-b-2 border-black'
              : 'text-gray-500'
          }`}
        >
          Type de paiement
        </button>
      </div>

      {/* Graphique Doughnut */}
      <div className="relative w-56 h-56 mx-auto mb-4">
        <Doughnut
          data={activeTab === 0 ? genderData : paymentData}
          options={options}
        />
      </div>

      {/* Légende personnalisée */}
      <div className="flex justify-center space-x-4 mt-4">
        {(activeTab === 0 ? genderData.labels : paymentData.labels).map(
          (label, index) => (
            <div key={index} className="flex items-center space-x-1">
              <div
                className="w-4 h-4 rounded-full"
                style={{
                  backgroundColor: (activeTab === 0 ? genderData : paymentData)
                    .datasets[0].backgroundColor[index],
                }}
              />
              <span className="text-sm text-gray-700">{label}</span>
            </div>
          ),
        )}
      </div>
    </div>
  );
};

export default GenderPaymentChart;
