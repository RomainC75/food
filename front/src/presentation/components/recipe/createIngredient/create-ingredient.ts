import { BaseUnit } from "@/modules/recipe/domain/baseUnit";
import { z } from "zod";

export const createIngredientSchema = z.object({
  name: z.string().min(2, "Name is required"),

  baseUnit: z.enum(Object.values(BaseUnit)),
});

export type CreateIngredientForm = z.infer<
  typeof createIngredientSchema
>;