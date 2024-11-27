// Typographie.stories.tsx
import H1 from './H1';
import H2 from './H2';
// import H3 from './H3';
// import H4 from './H4';
// import H5 from './H5';
// import TextBody from './TextBody';

export default {
  title: 'UI/Typographie', // Modifiez le titre selon vos besoins
  component: H1, // Composant principal
};

// Story pour H1
export const Heading1 = () => <H1 className="text-center">Ceci est un H1</H1>;

// Story pour H2
export const Heading2 = () => <H2 className="text-center">Ceci est un H2</H2>;

// // Story pour H3
// export const Heading3 = () => (
//   <H3 className="text-center">Ceci est un H3</H3>
// );

// // Story pour H4
// export const Heading4 = () => (
//   <H4 className="text-center">Ceci est un H4</H4>
// );

// // Story pour H5
// export const Heading5 = () => (
//   <H5 className="text-center">Ceci est un H5</H5>
// );

// // Story pour TextBody
// export const BodyText = () => (
//   <TextBody className="text-center">Ceci est un texte de corps.</TextBody>
// );
