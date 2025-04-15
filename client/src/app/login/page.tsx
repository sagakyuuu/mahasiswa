import { Form } from "@/components/ui/form";
import React from "react";
import { z } from "zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";

const formSchema = z.object({
  email: z.string().email({
    message: "email tidak sesuai",
  }),
  password: z.string().min(6, {
    message: "minimal 6 karakter",
  }),
});

export default function LoginPage() {
  const mutation = useMutation({
    mutationFn: async (e) => {
      e.preventDefault();
      try {
        const res = await fetch("http://localhost:5000/api/v1/auth/login", {});
      } catch (error) {}
    },
  });

  const form = useForm({});
  return (
    <Form>
      <form></form>
    </Form>
  );
}
