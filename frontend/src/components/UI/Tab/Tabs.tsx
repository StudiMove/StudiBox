import { useState } from 'react';

interface Tab {
  label: string;
  component: JSX.Element;
}

interface TabsProps {
  tabs: Tab[];
}

const Tabs = ({ tabs }: TabsProps) => {
  const [activeTab, setActiveTab] = useState(0);

  return (
    <div className="w-full">
      {/* Conteneur scrollable horizontalement pour la barre des onglets */}
      <div className="overflow-x-auto">
        <div className="flex items-center whitespace-nowrap">
          {tabs.map((tab, index) => (
            <button
              key={index}
              className={`text-lg pb-2 font-helvetica px-4 transition-all duration-300 ${
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
      </div>

      {/* Affichage du contenu de l'onglet actif */}
      <div className="mt-8">{tabs[activeTab].component}</div>
    </div>
  );
};

export default Tabs;
