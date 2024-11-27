import { useState, useEffect } from 'react';
import SearchBar from '../SearchBar/SearchBar';

interface Tab {
  label: string;
  dataLoader: () => Promise<any[]>;
}

interface TabsDataProps {
  tabs: Tab[];
  headers: string[];
  DataTableComponent: (props: {
    data: any[];
    headers: string[];
  }) => JSX.Element;
  mapData: (data: any[]) => any[];
  filterKey: string;
}

function TabsData({
  tabs,
  headers,
  DataTableComponent,
  mapData,
  filterKey,
}: TabsDataProps) {
  const [activeTab, setActiveTab] = useState(0);
  const [data, setData] = useState<any[]>([]);
  const [filteredData, setFilteredData] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [searchTerm, setSearchTerm] = useState('');

  const fetchData = async (dataLoader: () => Promise<any[]>) => {
    setLoading(true);
    try {
      const result = await dataLoader();
      const mappedData = mapData(result);
      setData(mappedData);
      setFilteredData(mappedData);
      setErrorMessage(null);
    } catch (error) {
      setErrorMessage('Erreur lors de la récupération des données.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData(tabs[activeTab].dataLoader);
  }, [activeTab]);

  useEffect(() => {
    const filtered = searchTerm
      ? data.filter((item) =>
          item[filterKey]?.toLowerCase().includes(searchTerm.toLowerCase()),
        )
      : data;
    setFilteredData(filtered);
  }, [searchTerm, data, filterKey]);

  return (
    <div className="w-full">
      <div className="py-8">
        <SearchBar onSearch={setSearchTerm} />
      </div>
      <div className="overflow-x-auto whitespace-nowrap flex items-center">
        {tabs.map((tab, index) => (
          <button
            key={index}
            className={`text-lg pb-2 px-4 transition-all duration-300 ${
              index === activeTab
                ? 'font-bold text-primary border-b-2 border-primary'
                : 'text-black'
            }`}
            onClick={() => setActiveTab(index)}
          >
            {tab.label}
          </button>
        ))}
      </div>
      <div className="mt-8">
        {loading ? (
          <div>Chargement des données...</div>
        ) : errorMessage ? (
          <div>{errorMessage}</div>
        ) : (
          <DataTableComponent data={filteredData} headers={headers} />
        )}
      </div>
    </div>
  );
}

export default TabsData;
