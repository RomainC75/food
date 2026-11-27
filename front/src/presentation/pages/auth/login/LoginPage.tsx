import { useForm } from "react-hook-form";
import { useShallow } from "zustand/shallow";
import { loginSchema } from "./login";
import { zodResolver } from "@hookform/resolvers/zod";
import { Box, Button, TextField, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import { LoginFormData } from "./login";
import { LoginError } from "@/presentation/state-manager/stores/auth.store";
import { nextPathFromLocation } from "@/presentation/components/auth/redirect";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import Link from "next/link";


const LOGIN_ERROR_MESSAGES: Record<LoginError, string> = {
  invalid_credentials: "Wrong email or password",
  unverified: "Account not verified, please check your emails",
  unknown: "Error trying to login, try again",
};

const LoginComponent = () => {
  const router = useRouter();
  const { loading, error, login } = useAuthStore(
    useShallow((state) => ({
      loading: state.loginPending,
      error: state.loginError,
      login: state.login,
    })),
  );

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
    mode: "onChange",
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (data: LoginFormData) => {
    const success = await login(data);
    if (success) {
      router.push(nextPathFromLocation());
    }
  };

  return (
    <div>
      LoginComponent
      <Box
        component="form"
        onSubmit={handleSubmit(onSubmit)}
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: 2,
          maxWidth: 400,
          mx: "auto",
        }}
      >
        <Typography variant="h5">Connect to your account</Typography>

        <TextField
          label="Email"
          type="text"
          {...register("email")}
          error={!!errors.email}
          helperText={errors.email?.message}
        />

        <TextField
          label="Password"
          type="password"
          {...register("password")}
          error={!!errors.password}
          helperText={errors.password?.message}
        />


        <Button type="submit" variant="contained" disabled={
          !!errors.email?.message ||
          !!errors.password?.message ||
          loading
        }>
          {loading ? "Login in..." : "Login"}
        </Button>
      </Box>
      <Link className="simple-link" href={"/auth/signup"}>No account yet ? Signup !</Link>
      {error && <div role="alert">{LOGIN_ERROR_MESSAGES[error]}</div>}
    </div>
  );
};

export default LoginComponent;
