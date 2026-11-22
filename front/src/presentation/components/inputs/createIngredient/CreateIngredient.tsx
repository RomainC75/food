"use client";

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";


import { Input } from "../../shadcn/input";
import { Button } from "../../shadcn/button";
import { CreateIngredientForm, createIngredientSchema } from "./create-ingredient";
import { useRecipeStore } from "@/app/provider";
import { BaseUnit } from "@/modules/recipe/domain/baseUnit";


export function CreateIngredientFormComponent() {
      const saveIngredient = useRecipeStore(
        (state) => state.saveIngredient
      );
  const form = useForm<CreateIngredientForm>({
    resolver: zodResolver(createIngredientSchema),

    defaultValues: {
      name: "",
      baseUnit: BaseUnit.Kg,
    },
  });

  const onSubmit = async (values: CreateIngredientForm) => {
    saveIngredient(values)
  };

  return (
    <form onSubmit={form.handleSubmit(onSubmit)}>
      <div>
        <label>Name</label>
        <Input {...form.register("name")} />
        {form.formState.errors.name && (
          <p>{form.formState.errors.name.message}</p>
        )}
      </div>

      <div>
        <label>Unit</label>
        <select {...form.register("baseUnit")}>
          <option value="kg">Kilogram</option>
          <option value="g">Gram</option>
          <option value="l">Liter</option>
          <option value="ml">Milliliter</option>
          <option value="unit">Unit</option>
        </select>
      </div>

      <Button type="submit">
        Create ingredient
      </Button>
    </form>
  );
}