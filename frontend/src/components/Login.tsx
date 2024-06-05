import { TextInput, Button, Group, Box, PasswordInput, LoadingOverlay, Modal } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useState } from "react";
import { Link } from "react-router-dom";

import CONFIG from "../config";

type LoginFormValues = {
  email: string;
  password: string;
};

function Login() {
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      password: (value) => (value.length >= 0 ? null : "Password is too short"),
    },
  });

  const doAuth = async (values: LoginFormValues) => {
    try {
      setLoading(true);
      const response = await fetch(`${CONFIG.backendURL}/api/user/auth`, {
        method: "POST",
        body: JSON.stringify(values),
        credentials: "include",
      });

      if (!response.ok) {
        let errorText = await response.text();
        throw new Error("Login failed: " + errorText);
      }

      setSuccess(true);

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  }


  return (
    <Box maw={340} mx="auto">
      <LoadingOverlay visible={loading} />

      {error &&
        <Modal opened={true} onClose={() => { setError(null) }} withCloseButton={true} title={error} />}

      <h2>Log in to your account</h2>
      <form onSubmit={form.onSubmit((values) => doAuth(values))}>
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
          <Button type="submit">Log in</Button>
        </Group>
      </form>

      <p>Don"t have an account? <Link to="/signup">Create an account</Link></p>

      {success && <p>Success!</p>}
    </Box>
  );
}

export default Login;