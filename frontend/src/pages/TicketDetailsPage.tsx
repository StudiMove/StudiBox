import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { getTicketDetails } from "../services/TicketService";
import { TicketDetailsResponse } from "../types/api/ticket/TicketResponse";
import { format, parseISO } from "date-fns";
import { fr } from "date-fns/locale"; // Pour le format français

const TicketDetailsPage = () => {
  const { uuid } = useParams<{ uuid: string }>();
  const [ticketDetails, setTicketDetails] = useState<TicketDetailsResponse | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!uuid) {
      setError("L'identifiant UUID est manquant.");
      setLoading(false);
      return;
    }

    const fetchDetails = async () => {
      try {
        const data = await getTicketDetails(uuid);
        setTicketDetails(data);
      } catch (err) {
        console.error("Erreur lors de la récupération des détails du ticket:", err);
        setError("Impossible de charger les détails du ticket.");
      } finally {
        setLoading(false);
      }
    };

    fetchDetails();
  }, [uuid]);

  if (loading) {
    return <p className="text-center text-lg">Chargement...</p>;
  }

  if (error) {
    return <p className="text-center text-red-500 text-lg">Erreur : {error}</p>;
  }

  if (!ticketDetails) {
    return <p className="text-center text-gray-600">Aucun détail disponible pour ce ticket.</p>;
  }

  // Vérification pour éviter les erreurs de type null/undefined
  const isTicketValid = ticketDetails.Status === "valid";
  const formattedDate =
    ticketDetails.Event && ticketDetails.Event.startDate
      ? format(parseISO(ticketDetails.Event.startDate), "dd MMMM yyyy", { locale: fr })
      : "Non définie";
  const formattedTime =
    ticketDetails.Event && ticketDetails.Event.startTime
      ? format(parseISO(ticketDetails.Event.startTime), "HH:mm", { locale: fr })
      : "Non définie";

  return (
    <div className="max-w-2xl mx-auto p-6">
      {/* En-tête avec état du billet */}
      <div
        className={`rounded-lg p-4 mb-6 ${
          isTicketValid ? "bg-green-100 border border-green-600" : "bg-red-100 border border-red-600"
        }`}
      >
        <h1 className={`text-xl font-semibold ${isTicketValid ? "text-green-600" : "text-red-600"}`}>
          {isTicketValid ? "Billet Valide" : "Billet Invalide"}
        </h1>
        <p className="mt-2 text-sm">
          {isTicketValid
            ? "Merci d'avoir utilisé votre billet ! Voici les détails de votre participation."
            : "Attention, ce billet n'est pas valide. Veuillez contacter notre support pour plus d'informations."}
        </p>
      </div>

      {/* Si le billet est valide, afficher les détails */}
      {isTicketValid ? (
        <>
          {/* Profil */}
          <div className="mb-6">
            <h2 className="text-lg font-semibold mb-2">Profil :</h2>
            <div className="border border-gray-300 rounded-lg p-4">
              <p>
                <span className="font-semibold">Nom :</span> {ticketDetails.User?.lastName || "Non défini"}
              </p>
              <p>
                <span className="font-semibold">Prénom :</span> {ticketDetails.User?.firstName || "Non défini"}
              </p>
              <p>
                <span className="font-semibold">Tarif :</span>{" "}
                {ticketDetails.Tarifs?.[0]?.Tarif?.title || "Non défini"}
              </p>
              <p>
                <span className="font-semibold">Âge :</span> {ticketDetails.User?.birthDate || "Non défini"}
              </p>
            </div>
          </div>
                  {/* Options */}
          <div className="mb-6">
            <h2 className="text-lg font-semibold mb-2">Options :</h2>
            {ticketDetails.Options && ticketDetails.Options.length > 0 ? (
              <div className="border border-gray-300 rounded-lg p-4">
                <ul className="list-disc pl-5">
                  {ticketDetails.Options.map((option, index) => (
                    <li key={index}>
                      <span className="font-semibold">{option.Option.title}</span>
                    </li>
                  ))}
                </ul>
              </div>
            ) : (
              <p className="text-gray-600">Aucune option sélectionnée.</p>
            )}
          </div>
          {/* Événement */}
          <div className="mb-6">
            <h2 className="text-lg font-semibold mb-2">Événement :</h2>
            <div className="border border-gray-300 rounded-lg p-4">
              <p>
                <span className="font-semibold">Nom :</span> {ticketDetails.Event?.title || "Non défini"}
              </p>
              <p>
                <span className="font-semibold">Date :</span> {formattedDate}
              </p>
              <p>
                <span className="font-semibold">Heure :</span> {formattedTime}
              </p>
            </div>
          </div>
        </>
      ) : (
        // Si le billet n'est pas valide, afficher un message d'erreur
        <div className="text-center text-gray-600">
          <p>Ce billet est invalide et ne peut pas être affiché.</p>
          <p>Veuillez contacter notre support pour assistance.</p>
        </div>
      )}
    </div>
  );
};

export default TicketDetailsPage;
