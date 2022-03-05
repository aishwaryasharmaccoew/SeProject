   export interface Item {
    id: string,
    background_image: string;
    name: string;
    released: string;
    description: string;
    genres: Array<Genre>;
    publishers: Array<Publishers>;
    ratings: Array<Rating>;
    screenshots: Array<Screenshots>;
  }
  
  export interface APIResponse<T> {
      results: Array<T>;
      
  }
  
  export interface Product {
    Id: string
    Name :        string
    Price :       Float32Array
    Category   :  string
    ImageUrl   :  string
    Product_url:  string
    Colors      : Array<Color>;
    Key_features : Array<Features>;
    Tools        :  Array<Tool>;
  }
  

  interface Color {
    name: string;
  }
  interface Features {
    name: string;
  }

  interface Tool {
    name: string;
  }


  interface Genre {
    name: string;
  }
  
  
  interface Publishers {
    name: string;
  }
  
  interface Rating {
    id: number;
    count: number;
    title: string;
  }
  
  interface Screenshots {
    image: string;
  }