import { Line } from 'react-chartjs-2';

interface LineChartProps {
  data: any;
  options: any;
}

const LineChart = ({ data, options }: LineChartProps) => {
  return <Line data={data} options={options} />;
};

export default LineChart;
