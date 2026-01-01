"use client";

import { GalleryVerticalEnd } from "lucide-react";

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
import axios from "axios";

export function SetPasswordForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const checkPasswordStrength = (password: string) => {
    const minLength = 8;
    const hasUpperCase = /[A-Z]/.test(password);
    const hasNumber = /\d/.test(password);
    const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password);

    return (
      password.length >= minLength &&
      hasUpperCase &&
      hasNumber &&
      hasSpecialChar
    );
  };

  const router = useRouter();

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    const form = event.target as HTMLFormElement;
    const password = (form.elements.namedItem("password") as HTMLInputElement)
      .value;
    const confirmPassword = (
      form.elements.namedItem("confirm-password") as HTMLInputElement
    ).value;

    if (password !== confirmPassword) {
      toast.error("Passwords do not match.");
      return;
    }

    if (!checkPasswordStrength(password)) {
      toast.error("Password does not meet the required strength criteria.");
      return;
    }

    const userEmail = localStorage.getItem("otp_email");
    if (!userEmail) {
      toast.error("Session expired. Please login again.");
      router.push("/login");
    }

    axios
      .post("http://localhost:8080/auth/set-password", {
        email: userEmail,
        password,
      })
      .then((response) => {
        toast.success(
          "Password set successfully! Login with your new password."
        );
        router.push("/login");
      })
      .catch((error) => {
        toast.error("An error occurred while setting the password.");
      });
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <form onSubmit={handleSubmit}>
        <FieldGroup>
          <div className="flex flex-col items-center gap-2 text-center">
            <a
              href="#"
              className="flex flex-col items-center gap-2 font-medium"
            >
              <div className="flex size-8 items-center justify-center rounded-md">
                <GalleryVerticalEnd className="size-6" />
              </div>
              <span className="sr-only">Omnicampus</span>
            </a>
            <h1 className="text-xl font-bold">Welcome to Omnicampus.</h1>
            <FieldDescription>
              Please enter your new password to continue.
              <ul className="mt-2 list-disc pl-5 text-left">
                <li>At least 8 characters</li>
                <li>At least one uppercase letter</li>
                <li>At least one number</li>
                <li>At least one special character</li>
              </ul>
            </FieldDescription>
          </div>
          <Field>
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <Input
              id="password"
              type="password"
              placeholder="Enter your new password"
              required
            />
          </Field>
          <Field>
            <FieldLabel htmlFor="confirm-password">Confirm Password</FieldLabel>
            <Input
              id="confirm-password"
              type="password"
              placeholder="Confirm your new password"
              required
            />
          </Field>
          <Field>
            <Button type="submit">Set Password</Button>
          </Field>
        </FieldGroup>
      </form>
    </div>
  );
}
