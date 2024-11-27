import upArrowIcon from '../Icons/arrowTrending/arrow-trending-up.svg'; // Icône flèche vers le haut
import downArrowIcon from '../Icons/arrowTrending/arrow-trending-down.svg'; // Icône flèche vers le bas

interface StatInfoBoxProps {
  title: string;
  value: number;
  percentage: number;
}

const StatInfoBox = ({ title, value, percentage }: StatInfoBoxProps) => {
  const radius = 20; // Rayon du cercle
  const circumference = 2 * Math.PI * radius; // Circonférence du cercle
  const strokeDashoffset = circumference - (percentage / 100) * circumference; // Calcul du pourcentage

  // Condition pour changer l'icône et la couleur de fond
  const isPositive = percentage > 33;
  const bgColor = isPositive ? 'bg-green-100' : 'bg-red-100';
  const icon = isPositive ? upArrowIcon : downArrowIcon;

  return (
    <div className="flex items-center justify-between p-4 bg-white rounded-xl shadow-md ">
      {/* Partie gauche avec l'icône et le titre */}
      <div className="flex items-center">
        <div className={`${bgColor} p-2 rounded-lg`}>
          {/* Icône qui change selon le pourcentage */}
          <img src={icon} alt="arrow-icon" className="h-6 w-6" />
        </div>
        <span className="ml-4 text-lg font-semibold">
          {title} : {value}
        </span>
      </div>

      {/* Partie droite avec le cercle pour le pourcentage */}
      <div className="relative">
        <svg width="50" height="50">
          {/* Cercle de progression (bleu) */}
          <circle
            cx="25"
            cy="25"
            r={radius}
            fill="none"
            stroke="#0F52BA"
            strokeWidth="2"
            strokeDasharray={circumference}
            strokeDashoffset={strokeDashoffset}
            strokeLinecap="round"
            transform="rotate(-90 25 25)" // Pour démarrer à 12h
          />
        </svg>
        {/* Texte du pourcentage au centre du cercle */}
        <div className="absolute inset-0 flex items-center justify-center text-blue-500 text-xs font-base">
          {percentage}%
        </div>
      </div>
    </div>
  );
};

export default StatInfoBox;
