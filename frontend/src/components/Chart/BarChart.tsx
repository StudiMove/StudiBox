import { Bar } from 'react-chartjs-2';

interface BarChartProps {
  data: any;
  options: any;
}

const BarChart = ({ data, options }: BarChartProps) => {
  return <Bar data={data} options={options} />;
};

export default BarChart;
