import { FormEventHandler } from "react";
import FormSubmit from "./components/FormSubmit";

interface FormProperties {
  className?: string;
  children?: any;
  onSubmit?: FormEventHandler<HTMLFormElement> | undefined;
}

const Form = ({ children, className, onSubmit }: FormProperties) => {
  return (
    <form className={className} onSubmit={onSubmit} action="">
      {children}
    </form>
  );
};

Form.Submit = FormSubmit;

export default Form;
