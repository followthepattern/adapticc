import { FormEvent } from "react";
import Form from "../../../components/headless/Form/Form";
import { useForm, SubmitHandler } from "react-hook-form";
import { gql, useMutation } from "@apollo/client";
import { login } from "../../../graphql/auth/login";
import { useUserStore } from "../../../utils/store";

type FormValues = {
  email: string;
  password: string;
};

const LoginForm = () => {
  const { register, handleSubmit } = useForm<FormValues>();
  const setJwt = useUserStore(state => state.setJwt);

  const [executeLogin, { data, loading, error }] = useMutation(
    gql(login)
  );

  const onSubmit: SubmitHandler<FormValues> = (formValues) => {
    executeLogin({
      variables: {
        email: formValues.email,
        password: formValues.password,
      },
    });
  };

  if (error) {
    return <div>{JSON.stringify(error)}</div>;
  }

  if (!data) {
    return (
      <Form className="max-w-md xs:w-full" onSubmit={handleSubmit(onSubmit)}>
        <div className="relative z-0 mb-6 w-full group">
          <input
            type="email"
            className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-gray-500 focus:outline-none focus:ring-0 focus:border-gray-600 peer"
            placeholder=" "
            required
            {...register("email")}
          />
          <label
            htmlFor="email"
            className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-gray-600 peer-focus:dark:text-gray-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          >
            Email address
          </label>
        </div>
        <div className="relative z-0 mb-6 w-full group">
          <input
            type="password"
            id="password"
            className="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-gray-500 focus:outline-none focus:ring-0 focus:border-gray-600 peer"
            placeholder=" "
            required
            {...register("password")}
          />
          <label
            htmlFor="password"
            className="absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-6 scale-75 top-3 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-gray-600 peer-focus:dark:text-gray-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6"
          >
            Password
          </label>
        </div>
        <div className="md:flex md:justify-center">
          <Form.Submit className="text-white bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm md:w-60 xs:w-full px-5 py-2.5 text-center dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800">
            {loading ? "Loading..." : "Login"}
          </Form.Submit>
        </div>
      </Form>
    );
  }

  setJwt(data.authentication.login.jwt)

  return <div>{JSON.stringify(data)}</div>;
};

export default LoginForm;
