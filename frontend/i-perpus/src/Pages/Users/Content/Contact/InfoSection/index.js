import React from "react";
import { useEffect, useState } from "react";
import { Button } from "../ButtonElement";
import {
  Column2,
  Img,
  ImgWrap,
  InfoContainer,
  InfoWrapper,
  InfoRow,
  Column1,
  TextWrapper,
  TopLine,
  Heading,
  Subtitle,
  BtnWrap,
} from "./InfoElements";

const InfoSection = ({
  lightBg,
  imgStart,
  topLine,
  lightText,
  headLine,
  darkText,
  description,
  buttonLabel,
  img,
  alt,
  id,
  primary,
  dark,
  dark2,
}) => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getArticles = async () => {
      const request = await fetch(
        "https://api.spaceflightnewsapi.net/v3/articles"
      );
      const response = await request.json();

      setArticles(response);
      setLoading(false);
    };
    getArticles();
  }, []);
  return (
    <>
      <section>
        {loading && <i>Loading articles ...</i>}
        {!loading && (
          <div>
            {articles.map((article) => {
              return (
                <InfoContainer lightBg={lightBg} id={id}>
                  <InfoWrapper>
                    <InfoRow imgStart={imgStart}>
                      <Column1>
                        <TextWrapper>
                          <TopLine>
                            {new Date(article.publishedAt).toLocaleDateString()}
                          </TopLine>
                          <Heading lightText={lightText}>
                            {article.title}
                          </Heading>
                          <Subtitle darkText={darkText}>
                            {article.summary}
                          </Subtitle>
                          <BtnWrap>
                            <Button
                              to={`/karya-tulis/${article.id}`}
                              smooth={true}
                              duration={500}
                              spy={true}
                              exact="true"
                              offset={-80}
                              primary={primary ? 1 : 0}
                              dark={dark ? 1 : 0}
                              dark2={dark2 ? 1 : 0}
                            >
                              Baca
                            </Button>
                          </BtnWrap>
                        </TextWrapper>
                      </Column1>
                      <Column2>
                        <ImgWrap>
                          <Img src={article.imageUrl} alt={alt} />
                        </ImgWrap>
                      </Column2>
                    </InfoRow>
                  </InfoWrapper>
                </InfoContainer>
              );
            })}
          </div>
        )}
      </section>
    </>
  );
};

export default InfoSection;
