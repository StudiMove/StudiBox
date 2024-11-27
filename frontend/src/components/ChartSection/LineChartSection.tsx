import LineChart from '../Chart/LineChart';

const LineChartSection = () => {
  const lineChartData = {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
    datasets: [
      {
        label: 'Recette totale (barre)',
        data: [100, 200, 150, 180, 220, 300], // Vérifie ces valeurs
        fill: false,
        backgroundColor: 'rgba(54, 162, 235, 0.2)',
        borderColor: 'rgba(54, 162, 235, 1)',
        tension: 0.4,
      },
      {
        label: 'Nombre de vue (ligne)',
        data: [50, 120, 80, 160, 90, 130], // Vérifie aussi ces valeurs
        fill: false,
        backgroundColor: 'rgba(255, 99, 132, 0.2)',
        borderColor: 'rgba(255, 99, 132, 1)',
        borderDash: [5, 5],
        tension: 0.4,
      },
    ],
  };

  //   const chartOptions = {
  //     responsive: true,
  //     plugins: {
  //       legend: {
  //         position: 'top',
  //       },
  //     },
  //   };
  const chartOptions = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
    },
    scales: {
      y: {
        beginAtZero: true, // Force l'axe Y à commencer à 0
        min: 0, // Valeur minimum (tu peux ajuster si nécessaire)
        max: 40000, // Valeur maximum basée sur tes données, ajuste selon tes besoins
      },
    },
  };

  return (
    <div className="col-span-2">
      <h2 className="text-xl font-bold mb-4">Statistiques des ventes</h2>
      <LineChart data={lineChartData} options={chartOptions} />
    </div>
  );
};

export default LineChartSection;
