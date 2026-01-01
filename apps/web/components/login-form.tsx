"use client";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
  FieldSeparator,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import axios, { AxiosError } from "axios";
import { useAuthStore } from "@/store/auth.store";

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"form">) {
  const router = useRouter();
  const { setUser } = useAuthStore.getState();

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const email = (form.elements.namedItem("email") as HTMLInputElement).value;
    const password = (form.elements.namedItem("password") as HTMLInputElement)
      .value;

    // OTP login flow
    if (password === "") {
      try {
        await axios.post(
          `${process.env.NEXT_PUBLIC_API_URL}/auth/request-otp`,
          { email }
        );

        localStorage.setItem("otp_email", email);
        router.push("/verify-otp");
      } catch (err: unknown) {
        if (axios.isAxiosError(err)) {
          const status = err.response?.status;
          const message = err.response?.data?.error;

          if (status === 400) {
            toast.error(
              "Account already onboarded. Please enter password to login."
            );
            return;
          }

          if (status === 401) {
            toast.error("Unauthorized request. Please try again.");
            return;
          }

          if (status === 500) {
            toast.error("Server error. Please try again later.");
            return;
          }

          toast.error(message || "Failed to send OTP");
        } else {
          toast.error("Something went wrong");
        }
      }
      return;
    }

    try {
      const req = await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/login`,
        {
          email,
          password,
        }
      );

      const token = req.data.token;

      document.cookie = `auth_token=${token}; path=/; max-age=86400`;

      const res = await axios.get(
        `${process.env.NEXT_PUBLIC_API_URL}/auth/me`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const data = res.data;

      setUser({
        id: data.ID,
        name: data.Name,
        email: data.Email,
        register_number: data.RegisterNumber,
        date_of_birth: data.DOB,
        phone: data.Phone,
      });

      toast.success("Login successful!");

      router.push("/");
    } catch (err: unknown) {
      if (axios.isAxiosError(err)) {
        const status = err.response?.status;
        const message = err.response?.data?.error;

        if (status === 400) {
          toast.error(
            "Account already onboarded. Please enter password to login."
          );
          return;
        }

        if (status === 401) {
          toast.error("Unauthorized request. Please try again.");
          return;
        }

        if (status === 500) {
          toast.error("Server error. Please try again later.");
          return;
        }

        toast.error(message || "Failed to send OTP");
      } else {
        toast.error("Something went wrong");
      }
    }
  };

  return (
    <form
      className={cn("flex flex-col gap-6", className)}
      {...props}
      onSubmit={onSubmit}
    >
      <FieldGroup>
        <div className="flex flex-col items-center gap-1 text-center">
          <h1 className="text-2xl font-bold">Login to your account</h1>
          <p className="text-muted-foreground text-sm text-balance">
            Enter your email below to login to your account
          </p>
        </div>
        <Field>
          <FieldLabel htmlFor="email">Email</FieldLabel>
          <Input id="email" type="email" placeholder="m@example.com" required />
        </Field>
        <Field>
          <div className="flex items-center">
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <a
              href="#"
              className="ml-auto text-sm underline-offset-4 hover:underline"
            >
              Forgot your password?
            </a>
          </div>
          <Input id="password" type="password" />
        </Field>
        <Field>
          <Button type="submit">Login</Button>
        </Field>
        <FieldSeparator>Or continue with</FieldSeparator>
        <Field>
          <Button variant="outline" type="button">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <path
                d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"
                fill="currentColor"
              />
            </svg>
            Login with Google
          </Button>
          <FieldDescription className="text-center">
            Don&apos;t have an account? Press login without entering a password!
          </FieldDescription>
        </Field>
      </FieldGroup>
    </form>
  );
}
