// import BarChart from '../../Chart/BarChart';
// import LineChartSection from '../../ChartSection/LineChartSection';
// const EventStatsComponent = () => {
//   const barChartData = {
//     labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
//     datasets: [
//       {
//         label: 'Ventes par mois',
//         data: [1000, 6000, 4000, 8000, 3000, 5000],
//         backgroundColor: 'rgba(54, 162, 235, 0.2)',
//         borderColor: 'rgba(54, 162, 235, 1)',
//         borderWidth: 1,
//       },
//     ],
//   };

//   // Options de base pour le BarChart
//   const chartOptions = {
//     responsive: true,
//     plugins: {
//       legend: {
//         position: 'top',
//       },
//       title: {
//         display: true,
//         text: 'Ventes par mois',
//       },
//     },
//   };

//   return (
//     <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
//       {/* Tester seulement le BarChart pour isoler le problème */}
//       <div className="col-span-2">
//         <LineChartSection />
//       </div>
//     </div>
//   );
// };

// export default EventStatsComponent;
import StatInfoBox from '../../Chart/StatInfoBox';
import SoldTickets from '../../Chart/SoldTickets';
import GenderPaymentChart from '../../Chart/GenderPaymentChart';
// import PersonInterested from '../../Chart/PersonInterested';
// import SalesStatsComponent from '../../Chart/SalesStatsComponent';
const EventStatsComponentDash = () => {
  const statInfoData = [
    { title: 'Étudiant', value: 23, percentage: 10 },
    { title: 'Externe', value: 57, percentage: 50 },
    { title: 'Âge moyen', value: 48, percentage: 48 },
  ];

  return (
    <>
      <div className="mb-8 flex gap-6">
        {/* <SalesStatsComponent /> */}
        <GenderPaymentChart />
        <SoldTickets
          stockPercentage={25}
          soldTickets={10000}
          totalRevenue={100000}
        />
      </div>
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
        {statInfoData.map((info, index) => (
          <StatInfoBox
            key={index}
            title={info.title}
            value={info.value}
            percentage={info.percentage}
          />
        ))}
      </div>

      <div className="mb-8 flex gap-6">{/* <PersonInterested /> */}</div>
    </>
  );
};

export default EventStatsComponentDash;
