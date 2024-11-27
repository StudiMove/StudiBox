import { Pie } from 'react-chartjs-2';

interface PieChartProps {
  data: any;
  options: any;
}

const PieChart = ({ data, options }: PieChartProps) => {
  return <Pie data={data} options={options} />;
};

export default PieChart;
