"use client";
import { useMutation } from "@tanstack/react-query";
import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { api } from "@/lib/api";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Link from "next/link";

const formShchema = z.object({
  email: z.string().email({ message: "Email tidak sesuai" }),
  password: z.string().min(6, { message: "Password minimal 6 karakter" }),
});

type FormData = z.infer<typeof formShchema>;

export default function Register() {
  const form = useForm<FormData>({
    resolver: zodResolver(formShchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const mutation = useMutation({
    mutationFn: async (data: FormData) => {
      const res = await api.post("/auth/register", data);
      console.error(res.data.error);
      return res.data;
    },
    onSuccess: (data) => {
      console.log("Berhasil dafter\n Token:" + data.token);
    },
    onError: (err: any) => {
      const message = err?.response?.data?.error ?? "Terjadi Kesalahan";
      console.log(message);
    },
  });

  const onSubmit = (data: FormData) => {
    console.log("Form Data:", data);

    mutation.mutate(data);
  };

  return (
    <div className="w-full min-h-screen flex justify-center items-center">
      <div className="min-w-xs border-2 p-4">
        <div className="mb-2 text-center">
          <h1 className="text-center text-3xl font-bold">Register</h1>
          <p>Masukan data anda untuk mendaftar</p>
        </div>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="space-y-6 max-w-sm"
          >
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      type="email"
                      placeholder="example@mail.com"
                      {...field}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input type="password" placeholder="******" {...field} />
                  </FormControl>
                </FormItem>
              )}
            />
            <Button type="submit" className="w-full">
              Daftar
            </Button>
          </form>
          <p className="text-center">
            Sudah punya akun?<Link href={"/login"}>Masuk</Link>
          </p>
        </Form>
      </div>
    </div>
  );
}
