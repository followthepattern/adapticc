import FormSubmit from "./components/FormSubmit";

interface FormProperties {
  className?: string;
  children?: any;
  onSubmit?: () => void;
}

const Form = ({ children, className }: FormProperties) => {
  return <form className={className}>{children}</form>;
};

Form.Submit = FormSubmit;

export default Form;
