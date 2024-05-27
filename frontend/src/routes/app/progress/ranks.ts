interface UserRanks {
    label: string,
    svgType: string,
    required_xp: Number
}

export let Ranks: Array<UserRanks> =  [
    {
        label: "Neuling",
        svgType: "Rank_1",
        required_xp: 0,
    },
    {
        label: "Anfänger",
        svgType: "Rank_2",
        required_xp: 360,
    },
    {
        label: "Aufstrebendes Mitglied",
        svgType: "Rank_3",
        required_xp: 720,
    },
    {
        label: "Mitglied",
        svgType: "Rank_4",
        required_xp: 1440,
    },
    {
        label:"Treues Mitglied",
        svgType:"Rank_5",
        required_xp: 2040,
    },
    {
        label: "Angesehenes Mitglied",
        svgType: "Rank_6",
        required_xp: 3000,
    },
    {
        label: "Wertvolles Mitglied",
        svgType: "Rank_7",
        required_xp: 3900,
    },
    {
        label:"Treuer Begleiter",
        svgType:"Rank_8",
        required_xp: 4900,
    },
    {
        label: "Partner",
        svgType: "Rank_9",
        required_xp: 6200,
    },
    {
        label:"Reisender",
        svgType:"Rank_10",
        required_xp: 7500,
    },
    {
        label: "Entdecker",
        svgType: "Rank_11",
        required_xp: 9800,
    },
    {
        label: "Meister der Selbstdisziplin",
        svgType: "Rank_12",
        required_xp: 12300,
    },
    {
        label:"Workaholic",
        svgType:"Rank_13",
        required_xp: 15100,
    },

    ]
