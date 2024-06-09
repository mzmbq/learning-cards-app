import { TextInput, Button, Group, Box, PasswordInput, LoadingOverlay } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Link, useNavigate } from "react-router-dom";

import CONFIG from "../config";
import { useState } from "react";

type SignUpFormValues = {
  email: string;
  password: string;
};

function SignUp() {
  const [success, setSuccess] = useState(false);
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

  const handleSubmit = async (values: SignUpFormValues) => {
    setIsLoading(true);

    try {
      const response = await fetch(`${CONFIG.backendURL}/api/user/create`, {
        method: "POST",
        body: JSON.stringify(values),
      });

      if (!response.ok) {
        throw new Error("Sign up failed");
      }

      setSuccess(true);

      navigate("/decks");

    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  // console.log(form.getInputProps("email"))

  return (
    <Box maw={340} mx="auto">
      <LoadingOverlay visible={isLoading} />

      {success && <p>Success!</p>}

      <h2>Create an account</h2>

      <form onSubmit={form.onSubmit(handleSubmit)}>
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