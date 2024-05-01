import { TextInput, Button, Group, Box, PasswordInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { Link } from 'react-router-dom';

function SignUp() {
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


  console.log(form.getInputProps("email"))

  return (
    <Box maw={340} mx="auto">
      <h2>Create an account</h2>
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
          <Button type="submit">Sign up</Button>
        </Group>
      </form>

      <p>Already have an account? <Link to="/login">Log in</Link></p>
    </Box>
  );
}

export default SignUp;