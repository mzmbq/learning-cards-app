import { TextInput, Button, Group, Box, PasswordInput, LoadingOverlay, Modal } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Link, useNavigate } from "react-router-dom";

import CONFIG from "../config";
import { useState } from "react";

type SignUpFormValues = {
  email: string;
  password: string;
};

function SignUp() {
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const navigate = useNavigate();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });

  const doSignup = async (values: SignUpFormValues) => {
    setIsLoading(true);

    try {
      const response = await fetch(`${CONFIG.backendURL}/api/user/create`, {
        method: "POST",
        body: JSON.stringify(values),
      });

      if (!response.ok) {
        if (response.status === 409) {
          throw new Error("A user with this email already exists");
        }
        throw new Error("Sign up failed");
      }


      navigate("/login");

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  // console.log(form.getInputProps("email"))

  return (
    <Box maw={340} mx="auto">
      <LoadingOverlay visible={isLoading} />

      {error &&
        <Modal opened={true} onClose={() => { setError(null); }} withCloseButton={true} title={error} />}


      <h2>Create an account</h2>

      <form onSubmit={form.onSubmit(doSignup)}>
        <TextInput
          label="Email Address"
          placeholder=""
          key={form.key("email")}
          {...form.getInputProps("email")}
        />

        <PasswordInput
          label="Password"
          placeholder=""
          key={form.key("password")}
          {...form.getInputProps("password")}
        />

        <Group justify="flex-end" mt="md">
          <Button type="submit">Sign up</Button>
        </Group>
      </form>

      <p>Already have an account? <Link to="/login">Log in</Link></p>
    </Box>
  );
}

export default SignUp;