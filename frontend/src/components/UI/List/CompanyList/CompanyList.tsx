// src/components/UI/List/CompanyList/CompanyList.tsx
import { Company } from '../../../../types/Table/DataTableTypes';
import DataTable from '../../Table/DataTable';

interface CompanyListProps {
  data: Company[];
  headers: string[];
  onRowClick: (id: number, name: string) => void;
}

function CompanyList({ data, headers, onRowClick }: CompanyListProps) {
  const handleClick = (company: Company) => {
    onRowClick(company.user_id, company.name);
  };

  return <DataTable data={data} headers={headers} onRowClick={handleClick} />;
}

export default CompanyList;
