import React from "react";
import { Accordion, AccordionItem, Button } from "@heroui/react";

export default function List({ data }) {
  return (
    <div className="max-w-[610px] pt-5">
      <Accordion defaultExpandedKeys={["2"]} variant="shadow">
      {data.map((item, idx) => (
        <AccordionItem key={idx} aria-label="Accordion 1" subtitle={item.album.Title} title={item.title}>
          <div className="flex flex-col">
            {item.lyrics}
            <div className="mt-3 flex gap-2">
              <Button color="primary" variant="flat">
                Сhange
              </Button>
              <Button color="primary" variant="light">
                Delete
              </Button>
            </div>
          </div>
        </AccordionItem>
      ))}
    </Accordion> 
    </div>
   
  );
}
