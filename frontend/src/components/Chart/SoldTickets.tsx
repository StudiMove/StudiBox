import { Doughnut } from 'react-chartjs-2';
import { Chart, ArcElement, Tooltip, Legend } from 'chart.js';
import ticketIcon from '../Icons/ticket.svg';
// Enregistrement des éléments du graphique
Chart.register(ArcElement, Tooltip, Legend);

interface SoldTicketsProps {
  stockPercentage: number; // Pourcentage du stock restant
  soldTickets: number; // Nombre de places vendues
  totalRevenue: number; // Recette totale
}

const SoldTickets = ({
  stockPercentage,
  soldTickets,
  totalRevenue,
}: SoldTicketsProps) => {
  const data = {
    labels: ['Stock restant', 'Stock utilisé'],
    datasets: [
      {
        data: [100 - stockPercentage, stockPercentage],
        backgroundColor: ['#d3d3d3', '#0F52BA'],
        hoverBackgroundColor: ['#c1c1c1', '#2980b9'],
        borderWidth: 0,
      },
    ],
  };

  const options = {
    cutout: '75%', // Crée l'effet d'anneau
    responsive: true,
    plugins: {
      legend: {
        display: false, // Masquer la légende pour ce graphique
      },
    },
  };

  return (
    <div className="bg-white rounded-xl p-6 shadow-md w-96 ">
      {' '}
      {/* Ajuste la taille globale */}
      <div className="flex flex-col items-center">
        {/* Titre avec icône */}
        <div className="flex items-center mb-4">
          <img src={ticketIcon} alt="icon" className="h-6 w-6 mr-2" />
          <h3 className="text-xl font-semibold text-black">Places Vendues</h3>
        </div>

        {/* Graphique Doughnut */}
        <div className="relative w-40 h-40 mb-4">
          <Doughnut data={data} options={options} />
          {/* Pourcentage au centre du graphique */}
          <div className="absolute inset-0 flex flex-col items-center justify-center text-center">
            <span className="text-sm text-gray-500">Stocks</span>
            <span className="text-xl font-bold">{stockPercentage}%</span>
          </div>
        </div>

        {/* Détails des places vendues et de la recette */}
        <div className="flex justify-between w-full text-gray-500 text-sm mt-4">
          <div className="text-left">
            <span className="block font-semibold text-black">
              Les places vendues :
            </span>
            <span>{soldTickets.toLocaleString()}</span>{' '}
            {/* Utilisation de toLocaleString pour le formatage des nombres */}
          </div>
          <div className="text-right">
            <span className="block font-semibold text-black">
              Recette totale :
            </span>
            <span>{totalRevenue.toLocaleString()}€</span>{' '}
            {/* Formatage du nombre */}
          </div>
        </div>
      </div>
    </div>
  );
};

export default SoldTickets;
