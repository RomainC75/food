"use client"
import { useRecipeStore } from '@/app/provider';
import React, { useEffect } from 'react'

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