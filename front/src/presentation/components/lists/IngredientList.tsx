"use client"
import { useRecipeStore } from '@/app/provider';
import React, { useEffect } from 'react'

export const NoIngredientFound = "no ingredient found, create a new one"

const IngredientList = () => {

    const ingredients = useRecipeStore((state) => state.ingredients);
    const loading = useRecipeStore((state) => state.loading);
    const error = useRecipeStore((state) => state.error);
    const fetchRecipes = useRecipeStore(
      (state) => state.fetchIngredients
    );
  
    useEffect(() => {
      console.log('-- useEffect')
      fetchRecipes();
  
    }, [fetchRecipes]);
  
    if (loading) return <p>Loading recipes...</p>;
  
    if (error) return <p>Error: {error}</p>;

    if (ingredients.length == 0) {
      return (
        <div>
          {NoIngredientFound}   
        </div>
      )
    }
  
  return (
    <div>
      <div>
        {ingredients.map((ing, index)=>
          <div key={index}>
            <div> id : {ing.Id}</div>
            <div> name : {ing.Name}</div>
            <div> unit : {ing.BaseUnit}</div>
          </div>
        )}
      </div>
    </div>
  )
}

export default IngredientList