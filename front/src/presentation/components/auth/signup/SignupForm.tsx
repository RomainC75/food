import React from "react";
import { useForm } from "react-hook-form";
import { useShallow } from "zustand/shallow";
import { SignupFormData, signupSchema } from "./signup";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import { Box, Button, TextField, Typography } from "@mui/material";

const SignupComponent = () => {
  const { loading, error, signup } = useAuthStore(
    useShallow((state) => ({
      loading: state.loading,
      error: state.error,
      signup: state.signup,
    })),
  );

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupFormData>({
    resolver: zodResolver(signupSchema),
    mode: "onChange",
    defaultValues: {
      email: "",
      password: "",
      confirmPassword: "",
      firstName: "",
      lastName: "",
    },
  });

  const onSubmit = async (data: SignupFormData) => {
    console.log("-> data : ", data);
    signup(data)
  };

  return (
    <div>
      SignupComponent
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
        <Typography variant="h5">Create an account</Typography>

        <TextField
          label="First name"
          {...register("firstName")}
          error={!!errors.firstName}
          helperText={errors.firstName?.message}
        />

        <TextField
          label="Last name"
          {...register("lastName")}
          error={!!errors.lastName}
          helperText={errors.lastName?.message}
        />

        <TextField
          label="Email"
          type="email"
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

        <TextField
          label="Confirm password"
          type="password"
          {...register("confirmPassword")}
          error={!!errors.confirmPassword}
          helperText={errors.confirmPassword?.message}
        />

        <Button type="submit" variant="contained" disabled={
          !!errors.firstName?.message ||
          !!errors.lastName?.message ||
          !!errors.email?.message ||
          !!errors.password?.message ||
          !!errors.confirmPassword?.message ||
          loading
        }>
          {loading ? "Creating account..." : "Sign up"}
        </Button>
      </Box>
    </div>
  );
};

export default SignupComponent;
