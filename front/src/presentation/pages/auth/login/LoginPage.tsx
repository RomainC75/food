import React from "react";
import { useForm } from "react-hook-form";
import { useShallow } from "zustand/shallow";
import { loginSchema } from "./login";
import { zodResolver } from "@hookform/resolvers/zod";
import { Box, Button, TextField, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import { LoginFormData } from "./login";
import { ErrorUserUnverified } from "@/modules/auth/infra/in-mem-users";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";


const LoginComponent = () => {
  const router = useRouter();
  const { loading, error, login } = useAuthStore(
    useShallow((state) => ({
      loading: state.loading,
      error: state.error,
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
    if (success){
      router.push("/")
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
      <div>{error}</div>
      {error && error !== ErrorUserUnverified &&  <div>Error trying to login, try again</div>}
      {error === ErrorUserUnverified && <div>check you emails</div>}
    </div>
  );
};

export default LoginComponent;
