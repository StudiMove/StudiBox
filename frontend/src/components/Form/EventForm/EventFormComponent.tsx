import { useState, useEffect } from 'react';
import {
  EventForm,
  EventData,
  EventLocation,
  EventDescription,
  EventOption,
  EventTarif,
  EventUpdateData,
} from '../../../types/api/event/EventTypes';
import { EventCreateData } from '../../../types/api/event/EventCreateData';
import EventToggleBlock from './EventToggleBlock';
import GeneralInfo from './GeneralInfo';
import DescriptionBlock from './DescriptionBlock';
import TicketingBlock from './TicketingBlock';
import TarifsBlock from './TarifsBlock';
import OptionsBlock from './OptionsBlock';
import LocationBlock from './LocationBlock';
import SaveButton from '../../UI/Button/SaveButton';
import ImageUpload from './ImageUpload';
import {
  getEventById,
  createEvent,
  getAllEventTags,
  getAllEventCategories,
  updateEvent,
  uploadEventImages,
} from '../../../services/EventService';
import { useIsMobile } from '../../../hooks/useMediaQuery';
interface EventFormComponentProps {
  isUpdate?: boolean;
  targetEventId?: number;
}

const EventFormComponent = ({
  isUpdate = false,
  targetEventId,
}: EventFormComponentProps) => {
  const [tagOptions, setTagOptions] = useState<
    { value: string; label: string }[]
  >([]);
  const [categoryOptions, setCategoryOptions] = useState<
    { value: string; label: string }[]
  >([]);
  const isMobile = useIsMobile();
  const [images, setImages] = useState<File[]>([]); // Stocker les images sélectionnées

  const [formData, setFormData] = useState<EventForm>({
    isOnline: false,
    isPublic: false,
    title: '',
    subtitle: '',
    categories: [],
    tags: [],
    startDate: null,
    startTime: null,
    endDate: null,
    endTime: null,
    description: [],
    options: [],
    tarifs: [],
    location: {
      address: '',
      city: '',
      region: '',
      postcode: '',
      country: '',
    },
    images: [],
    use_studibox: false,
    ticketPrice: 0,
    ticketStock: 0,
    externalTicketingUrl: '',
    video_url: '',
  });

  // Récupérer les tags et catégories
  useEffect(() => {
    const fetchTags = async () => {
      try {
        const tags = await getAllEventTags();
        setTagOptions(tags.map((tag) => ({ value: tag, label: tag })));
      } catch (error) {
        console.error('Erreur lors de la récupération des tags:', error);
      }
    };

    const fetchCategories = async () => {
      try {
        const categories = await getAllEventCategories();
        setCategoryOptions(
          categories.map((category) => ({ value: category, label: category })),
        );
      } catch (error) {
        console.error('Erreur lors de la récupération des catégories:', error);
      }
    };

    fetchTags();
    fetchCategories();
  }, []);

  const handleImageUpload = async (images: File[]): Promise<string[]> => {
    try {
      const uploadedUrls = await uploadEventImages(images);
      const formattedUrls = uploadedUrls.map((url) => url.toString());
      console.log(formattedUrls);

      // Conserver les anciennes images et ajouter les nouvelles en tant que tableau
      const updatedImages = [...formData.images, ...formattedUrls];
      updateField('images', updatedImages); // Mettre à jour formData.images en tant que tableau

      return updatedImages;
    } catch (error) {
      console.error("Erreur lors de l'upload des images:", error);
      return [];
    }
  };

  const transformEventDataToForm = (event: EventData): EventForm => ({
    isOnline: event.is_online,
    isPublic: event.is_visible ?? false,
    title: event.title,
    subtitle: event.subtitle,
    categories: event.categories,
    tags: event.tags,
    startDate: event.start_date ? new Date(event.start_date) : null,
    startTime: event.start_time ? new Date(event.start_time) : null,
    endDate: event.end_date ? new Date(event.end_date) : null,
    endTime: event.end_time ? new Date(event.end_time) : null,
    description: event.descriptions,
    options: event.options,
    tarifs: event.tarifs,
    location: {
      address: event.address,
      city: event.city,
      region: event.region,
      postcode: event.postcode,
      country: event.country,
    },
    // images: event.image || [],
    images: event.image_urls ? event.image_urls : [], // Conversion de image_urls en tableau

    use_studibox: event.use_studibox ?? false,
    ticketPrice: event.ticket_price,
    ticketStock: event.ticket_stock,
    externalTicketingUrl: event.external_ticketing_url || '',
    video_url: event.video_url || '',
  });

  const updateField = <K extends keyof EventForm>(
    field: K,
    value: EventForm[K],
  ) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const transformFormDataToUpdateData = (): EventUpdateData => ({
    event_id: targetEventId!,
    title: formData.title,
    subtitle: formData.subtitle,
    start_date: formData.startDate?.toISOString() || '',
    end_date: formData.endDate?.toISOString() || '',
    start_time: formData.startTime?.toISOString() || '',
    end_time: formData.endTime?.toISOString() || '',
    is_online: formData.isOnline,
    is_visible: formData.isPublic,
    use_studibox: formData.use_studibox,
    ticket_price: formData.ticketPrice,
    ticket_stock: formData.ticketStock,
    location: formData.location,
    options: formData.options,
    tarifs: formData.tarifs,
    descriptions: formData.description,
    video_url: formData.video_url,
    images: formData.images,
    tags: formData.tags,
    categories: formData.categories,
    // Ajout de l'URL de la vidéo
  });
  const updateVideoUrl = (url: string) => {
    updateField('video_url', url);
  };
  const handleFormSubmit = async () => {
    try {
      const uploadedImageUrls =
        images.length > 0 ? await handleImageUpload(images) : [];
      const allImageUrls = [...formData.images, ...uploadedImageUrls];

      if (isUpdate && targetEventId) {
        const apiPayload: EventUpdateData = {
          ...transformFormDataToUpdateData(),
          images: allImageUrls,
        };
        await updateEvent(apiPayload);
        console.log('Event updated successfully');
      } else {
        const singleStringUrls = allImageUrls.join(', ');
        const apiPayload: EventCreateData = {
          ...transformFormDataToUpdateData(),
          images: singleStringUrls,
        };
        await createEvent(apiPayload);
        console.log('Event created successfully');
      }
    } catch (error) {
      console.error("Erreur lors de la soumission de l'événement:", error);
    }
  };

  useEffect(() => {
    if (isUpdate && targetEventId) {
      (async () => {
        try {
          const event = await getEventById(targetEventId);
          console.log(
            "Données de l'événement récupérées avant transformation:",
            event,
          );

          setFormData(transformEventDataToForm(event));
        } catch (error) {
          console.error("Erreur lors du chargement de l'événement :", error);
        }
      })();
    }
  }, [isUpdate, targetEventId]);
  return (
    <div className={`${isMobile && !isUpdate && !targetEventId ? 'mx-8' : ''}`}>
      <EventToggleBlock formData={formData} updateField={updateField} />
      <GeneralInfo
        formData={formData}
        updateField={updateField}
        categoryOptions={categoryOptions}
        onCategoryChange={(selected) =>
          updateField(
            'categories',
            selected.map((opt) => opt.value),
          )
        }
        tagOptions={tagOptions}
        onTagChange={(selected) =>
          updateField(
            'tags',
            selected.map((opt) => opt.value),
          )
        }
      />
      <DescriptionBlock
        descriptions={formData.description}
        updateField={(descriptions: EventDescription[]) =>
          updateField('description', descriptions)
        }
      />
      <TicketingBlock
        use_studibox={formData.use_studibox}
        ticketPrice={formData.ticketPrice}
        ticketStock={formData.ticketStock}
        externalTicketingUrl={formData.externalTicketingUrl}
        updateField={updateField}
      />
      {formData.use_studibox && (
        <>
          <TarifsBlock
            tarifs={formData.tarifs}
            updateField={(tarifs: EventTarif[]) =>
              updateField('tarifs', tarifs)
            }
          />
          <OptionsBlock
            options={formData.options}
            updateField={(options: EventOption[]) =>
              updateField('options', options)
            }
          />
        </>
      )}
      <LocationBlock
        location={formData.location}
        updateField={(location: EventLocation) =>
          updateField('location', location)
        }
      />
      <ImageUpload
        images={images}
        updateField={setImages}
        uploadImages={handleImageUpload}
        video_url={formData.video_url} // Passe video_url à ImageUpload
        updateVideoUrl={updateVideoUrl} // Passe updateVideoUrl pour modifier video_url
      />

      <div className="fixed right-4 bottom-4">
        <SaveButton onClick={handleFormSubmit} />
      </div>
    </div>
  );
};

export default EventFormComponent;
