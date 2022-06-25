import React from "react";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Button } from "../../ButtonElement";
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
  Embed,
  ReadBook,
} from "./InfoElements";

const InfoSectionDetail = ({
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
  const params = useParams();
  const [article, setArticle] = useState({});
  const [loading, setLoading] = useState(true);
  const [notFound, setNotFound] = useState(false);

  useEffect(() => {
    const getArticles = async () => {
      const request = await fetch(
        `https://api.spaceflightnewsapi.net/v3/articles/${params.id}`
      );

      if (!request.ok) {
        return setNotFound(true);
      }

      const response = await request.json();

      setArticle(response);
      setLoading(false);
    };
    getArticles();
  }, [params]);

  if (notFound) {
    return <h1>Article Tidak Ditemukan :( </h1>;
  }

  return (
    <>
      <InfoContainer lightBg={lightBg} id={id}>
        <InfoWrapper>
          <InfoRow imgStart={imgStart}>
            <Column1>
              <TextWrapper>
                <TopLine>
                  {new Date(article.publishedAt).toLocaleDateString()}
                </TopLine>
                <Heading lightText={lightText}>{article.title}</Heading>
                <Subtitle darkText={darkText}>{article.summary}</Subtitle>
              </TextWrapper>
            </Column1>
            <Column2>
              <ImgWrap>
                <Img src={article.imageUrl} alt={alt} />
              </ImgWrap>
            </Column2>
          </InfoRow>
        </InfoWrapper>
        <ReadBook>Baca Buku</ReadBook>
        <Embed src="https://documentcloud.adobe.com/view-sdk-demo/PDFs/Bodea Brochure.pdf"></Embed>
      </InfoContainer>
      );
    </>
  );
};

export default InfoSectionDetail;
