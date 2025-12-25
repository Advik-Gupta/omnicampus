"use client";

import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp";

export function OTPForm(props: React.ComponentProps<typeof Card>) {
  const router = useRouter();
  const [otp, setOtp] = useState("");
  const [loading, setLoading] = useState(false);

  const handleVerify = async (e: React.FormEvent) => {
    e.preventDefault();

    const email = localStorage.getItem("otp_email");

    if (!email) {
      alert("Session expired. Please login again.");
      router.push("/login");
      return;
    }

    if (otp.length !== 6) {
      alert("Please enter a valid 6-digit OTP");
      return;
    }

    try {
      setLoading(true);

      const res = await axios.post("http://localhost:8080/auth/verify-otp", {
        email,
        otp,
      });

      const token = res.data.token;

      // Store JWT in cookie
      document.cookie = `token=${token}; path=/; max-age=${
        60 * 60 * 24 * 7
      }; SameSite=Lax`;

      localStorage.removeItem("otp_email");

      router.push("/");
    } catch (err: any) {
      alert(err.response?.data?.message || "Invalid OTP");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card {...props}>
      <CardHeader>
        <CardTitle>Enter verification code</CardTitle>
        <CardDescription>We sent a 6-digit code to your email.</CardDescription>
      </CardHeader>

      <CardContent>
        <form onSubmit={handleVerify}>
          <FieldGroup>
            <Field>
              <FieldLabel htmlFor="otp">Verification code</FieldLabel>

              <InputOTP
                id="otp"
                maxLength={6}
                value={otp}
                onChange={setOtp}
                required
              >
                <InputOTPGroup className="gap-2.5 *:data-[slot=input-otp-slot]:rounded-md *:data-[slot=input-otp-slot]:border">
                  <InputOTPSlot index={0} />
                  <InputOTPSlot index={1} />
                  <InputOTPSlot index={2} />
                  <InputOTPSlot index={3} />
                  <InputOTPSlot index={4} />
                  <InputOTPSlot index={5} />
                </InputOTPGroup>
              </InputOTP>

              <FieldDescription>
                Enter the 6-digit code sent to your email.
              </FieldDescription>
            </Field>

            <FieldGroup>
              <Button type="submit" disabled={loading}>
                {loading ? "Verifying..." : "Verify"}
              </Button>

              <FieldDescription className="text-center">
                Didn&apos;t receive the code?{" "}
                <button
                  type="button"
                  className="underline"
                  onClick={() => router.push("/login")}
                >
                  Resend
                </button>
              </FieldDescription>
            </FieldGroup>
          </FieldGroup>
        </form>
      </CardContent>
    </Card>
  );
}
