import { TextInput, Button, Group, Box, PasswordInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { Link } from 'react-router-dom';


function Login() {
  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });



  return (
    <Box maw={340} mx="auto">
      <h2>Sign in to your account</h2>
      <form onSubmit={form.onSubmit((values) => console.log(values))}>
        <TextInput
          label="Email Address"
          placeholder=""
          key={form.key('email')}
          {...form.getInputProps('email')}
        />

        <PasswordInput
          label="Password"
          placeholder=""
          key={form.key('password')}
          {...form.getInputProps('password')}
        />

        <Group justify="flex-end" mt="md">
          <Button type="submit">Log in</Button>
        </Group>
      </form>

      <p>Don't have an account? <Link to="/signup">Create an account</Link></p>
    </Box>
  );
}

export default Login;