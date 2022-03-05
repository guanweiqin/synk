import React, { useContext } from "react";
import styled from "styled-components";
import { showPhoneQrcodeDialog } from "../../pages/home/components";
import { AppContext } from "../../shared/app_context";

export const ShowQrcodeModal = () => {
	const context = useContext(AppContext);
    const onClick = async (e) => {
    	e.preventDefault();
	    showPhoneQrcodeDialog({
	      context,content: (addr) => addr && `http://${addr}:27149/static/`
	    });
	};

    return (
    	<Atext href="#!" onClick={onClick}>使用手机上传</Atext>
	);
};
export const Atext = styled.a`
  text-align: center;
`;